package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"runtime"
	"time"

	"github.com/ckeyer/logrus"
	"github.com/funxdata/landlady/global"
	pb "github.com/funxdata/landlady/proto"
	"github.com/funxdata/landlady/task"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func init() {
	rootCmd.AddCommand(TaskCmd())
}

func TaskCmd() *cobra.Command {
	var (
		addr          string
		redisEndpoint string
		redisDB       int
		redisAuth     string
	)

	cmd := &cobra.Command{
		Use:   "task",
		Short: "启动任务管理中心",
		Run: func(cmd *cobra.Command, args []string) {
			redisOpts := []redis.DialOption{
				redis.DialDatabase(redisDB),
			}
			if redisAuth != "" {
				redisOpts = append(redisOpts, redis.DialPassword(redisAuth))
			}

			conn, err := redis.Dial("tcp", redisEndpoint, redisOpts...)
			if err != nil {
				logrus.Error(err)
				return
			}

			lis, err := net.Listen("tcp", addr)
			if err != nil {
				logrus.Error(err)
				return
			}

			s := grpc.NewServer(Logger(), Recovery())

			pb.RegisterTasksServer(s, task.NewTasks(conn))

			logrus.Infof("server start at %s", addr)
			if err := s.Serve(lis); err != nil {
				logrus.Fatal("Failed to host run: %v", err)
			}
		},
	}

	cmd.Flags().StringVarP(&addr, "grpc-addr", "s", ":6666", "grpc server address.")
	cmd.Flags().StringVar(&redisEndpoint, "redis-endpoint", "127.0.0.1:6379", "redis endpoint.")
	cmd.Flags().IntVar(&redisDB, "redis-db", 0, "redis db.")
	cmd.Flags().StringVar(&redisAuth, "redis-auth", "", "redis auth.")

	return cmd
}

func Logger() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		ret, err := handler(ctx, req)
		consume := time.Now().Sub(start)

		fds := logrus.Fields{
			"method":  info.FullMethod,
			"consume": consume.String(),
		}
		if err != nil {
			fds["error"] = err
			logrus.WithFields(fds).Error("request over.")
		} else {
			logrus.WithFields(fds).Info("request over.")
		}

		return ret, err
	})
}

func Recovery() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				out := stack(2)
				msg := fmt.Sprintf("[Recovery] panic recovered: %s\n%s", r, out.String())

				logrus.Error(msg)
				if global.IsDebug() {
					err = grpc.Errorf(codes.Internal, "%s", msg)
				} else {
					err = grpc.Errorf(codes.Internal, "%s", r)
				}
			}
		}()

		return handler(ctx, req)
	})
}

func stack(skip int) *bytes.Buffer {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}

		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return []byte("???")
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return []byte("???")
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, []byte("/")); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, []byte(".")); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, []byte("·"), []byte("."), -1)
	return name
}
