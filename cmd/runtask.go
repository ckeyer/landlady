package cmd

import (
	"net/http"
	"time"

	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/rpc"
	pb "github.com/funxdata/landlady/proto"
	"github.com/funxdata/landlady/targets/zufang58xian"
	"github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func init() {
	rootCmd.AddCommand(AddTaskCmd())
}

func AddTaskCmd() *cobra.Command {
	var (
		serverEndpoint string
		ipfsEndpoint   string
	)

	cmd := &cobra.Command{
		Use:     "runtask",
		Aliases: []string{"runt"},
		Short:   "启动URL任务",
		Run: func(cmd *cobra.Command, args []string) {

			runner := zufang58xian.New()

			cli := pb.NewPagesClient(rpc.Dial(serverEndpoint))
			ipfscli := shell.NewShell("ss.ckeyer.com:55001")

			count, err := runner.PageCount(http.DefaultClient)
			if err != nil {
				logrus.Fatalf("get counter failed, %s", err)
			}

			httpclient := http.DefaultClient
			for i := 0; i < count; i++ {
				urls, err := runner.ScanURLs(httpclient, i+1)
				if err != nil {
					logrus.Errorf("get urls failed at index %s", i+1)
				}
				for _, v := range urls {
					req, _ := http.NewRequest("GET", v, nil)
					req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
					info, body, err := runner.DownloadPage(httpclient, req)
					if err != nil {
						logrus.Errorf("download page for %s failed, %s", v, err)
						time.Sleep(time.Second)
						continue
					}
					hash, err := ipfscli.Add(body)
					if err != nil {
						logrus.Errorf("save body to ipfs failed, %s", err)
					}
					page := &pb.Page{
						Url:      info,
						Status:   pb.PageInfoStatus_Downloaded,
						HandleAt: time.Now(),
						Hash:     hash,
					}
					cli.Save(context.Background(), page)
				}
			}
			// if len(args) < 2 {
			// 	logrus.Fatal("required project name or URL args.")
			// }
			// prjName := args[0]
			// tasks := make([]*pb.Task, 0, len(args)-1)
			// for _, u := range args[1:] {
			// 	if _, err := url.Parse(u); err != nil {
			// 		logrus.Fatalf("invalid URL %s, %s", u, err)
			// 	}
			// 	tasks = append(tasks, &pb.Task{ProjectName: prjName, Url: u})
			// }

			// _, err := cli.PushTasks(context.Background(), &pb.TaskList{Items: tasks})
			// if err != nil {
			// 	logrus.Errorf("add tasks failed", err)
			// 	return
			// }
			// logrus.Infof("add tasks %v successful", len(tasks))
		},
	}

	cmd.Flags().StringVarP(&serverEndpoint, "server-addr", "s", "127.0.0.1:6666", "grpc server address.")
	cmd.Flags().StringVar(&ipfsEndpoint, "ipfs-endpoint", "ss.ckeyer.com:55001", "ipfs server endpoint")

	return cmd
}
