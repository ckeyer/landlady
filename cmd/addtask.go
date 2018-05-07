package cmd

import (
	"net/url"

	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/rpc"
	pb "github.com/funxdata/landlady/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func init() {
	rootCmd.AddCommand(AddTaskCmd())
}

func AddTaskCmd() *cobra.Command {
	var (
		serverEndpoint string
	)

	cmd := &cobra.Command{
		Use:     "addtask",
		Aliases: []string{"addt"},
		Short:   "添加URL任务",
		Run: func(cmd *cobra.Command, args []string) {
			cli := pb.NewTasksClient(rpc.Dial(serverEndpoint))

			if len(args) < 2 {
				logrus.Fatal("required project name or URL args.")
			}
			prjName := args[0]
			tasks := make([]*pb.Task, 0, len(args)-1)
			for _, u := range args[1:] {
				if _, err := url.Parse(u); err != nil {
					logrus.Fatalf("invalid URL %s, %s", u, err)
				}
				tasks = append(tasks, &pb.Task{ProjectName: prjName, Url: u})
			}

			_, err := cli.AddTasks(context.Background(), &pb.TaskList{Items: tasks})
			if err != nil {
				logrus.Errorf("add tasks failed", err)
				return
			}
			logrus.Infof("add tasks %v successful", len(tasks))
		},
	}

	cmd.Flags().StringVarP(&serverEndpoint, "server-addr", "s", "127.0.0.1:6666", "grpc server address.")

	return cmd
}
