package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/rpc"
	"github.com/funxdata/commons/rpc/middleware"
	pb "github.com/funxdata/landlady/proto"
	"github.com/funxdata/landlady/task"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	redis "gopkg.in/redis.v5"
)

func init() {
	rootCmd.AddCommand(TaskCmd())
}

func TaskCmd() *cobra.Command {
	var (
		addr     string
		redisOpt = redis.Options{}
	)

	cmd := &cobra.Command{
		Use:   "task",
		Short: "启动任务管理中心",
		Run: func(cmd *cobra.Command, args []string) {
			rcli := redis.NewClient(&redisOpt)
			if err := rcli.Ping().Err(); err != nil {
				logrus.Fatalf("connect redis failed, %s", err)
			}

			s := grpc.NewServer(middleware.Logger(), middleware.Recovery())
			pb.RegisterTasksServer(s, task.NewTasks(rcli))

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

	return cmd
}
