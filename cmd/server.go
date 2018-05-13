package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/rpc"
	"github.com/funxdata/commons/rpc/middleware"
	"github.com/funxdata/landlady/pkg/tasks"
	pb "github.com/funxdata/landlady/proto"
	grpc_md "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

func init() {
	rootCmd.AddCommand(TaskCmd())
}

func TaskCmd() *cobra.Command {
	var (
		addr     string
		redisOpt = redis.Options{}
		mgoUrl   string
	)

	cmd := &cobra.Command{
		Use:   "server",
		Short: "启动任务管理中心",
		Run: func(cmd *cobra.Command, args []string) {
			rcli := redis.NewClient(&redisOpt)
			if err := rcli.Ping().Err(); err != nil {
				logrus.Fatalf("connect redis failed, %s", err)
			}

			mgoss, err := mgo.Dial(mgoUrl)
			if err != nil {
				logrus.Fatalf("connect mongodb failed, %s", err)
			}

			s := grpc.NewServer(grpc_md.WithUnaryServerChain(
				middleware.Logger(),
				middleware.Recovery(),
			))
			pb.RegisterTasksServer(s, tasks.NewTasks(rcli))

			logrus.Infof("server starting at %s", addr)
			if err := rpc.ServeTCP(s, addr); err != nil {
				logrus.Fatalf("Failed to host run: %v", err)
			}
		},
	}

	cmd.Flags().StringVarP(&addr, "grpc-addr", "s", ":6666", "grpc server address.")
	cmd.Flags().StringVar(&redisOpt.Addr, "redis-endpoint", "127.0.0.1:6379", "redis endpoint.")
	cmd.Flags().IntVar(&redisOpt.DB, "redis-db", 0, "redis db.")
	cmd.Flags().StringVar(&redisOpt.Password, "redis-auth", "", "redis auth.")
	cmd.Flags().StringVar(&mgoUrl, "mongodb-url", "mongodb://localhost:40001/zufang", "mongodb URL.")

	return cmd
}
