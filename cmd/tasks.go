package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/rpc"
	"github.com/funxdata/commons/rpc/middleware"
	pb "github.com/funxdata/landlady/proto"
	"github.com/funxdata/landlady/task"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

			s := grpc.NewServer(middleware.Logger(), middleware.Recovery())
			pb.RegisterTasksServer(s, task.NewTasks(conn))

			logrus.Infof("server starting at %s", addr)
			if err := rpc.ServeTCP(s, addr); err != nil {
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
