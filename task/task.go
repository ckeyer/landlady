package task

import (
	"sync"

	pb "github.com/funxdata/landlady/proto"
	"github.com/gogo/protobuf/types"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/context"
)

var _ pb.TasksServer = (*TaskController)(nil)

type TaskController struct {
	sync.Mutex
	redis.Conn
}

func NewTasks(rconn redis.Conn) *TaskController {
	return &TaskController{
		Conn: rconn,
	}
}

func (t *TaskController) NewProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	return nil, nil
}

func (t *TaskController) GetProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	return nil, nil
}

func (t *TaskController) AddTasks(ctx context.Context, in *pb.TaskList) (*pb.Task, error) {
	return nil, nil
}

func (t *TaskController) RequestTasks(ctx context.Context, in *pb.RequestTaskOption) (*pb.TaskList, error) {
	return nil, nil
}

func (t *TaskController) CompleteTask(ctx context.Context, in *pb.TaskResult) (*types.Empty, error) {
	return nil, nil
}
