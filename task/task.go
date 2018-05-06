package task

import (
	"strings"
	"sync"

	"github.com/funxdata/commons/gerr"
	pb "github.com/funxdata/landlady/proto"
	"github.com/gogo/protobuf/types"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/context"
)

const (
	rKeySep            = ":"
	rKeyPreTaskProject = "task_project"
)

var _ pb.TasksServer = (*TaskController)(nil)

type TaskController struct {
	sync.Mutex
	*redisClient
}

func NewTasks(rconn redis.Conn) *TaskController {
	return &TaskController{
		redisClient: newRedis(rconn),
	}
}

func (t *TaskController) NewProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	err := t.setKV(t.projectKey(in), in, false)

	if err != nil {
		return nil, gerr.Internal("set task project %s failed, %s", in.Name, err)
	}
	return in, nil
}

func (t *TaskController) GetProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	ret := &pb.TaskProject{}
	err := t.getKV(t.projectKey(in), ret)
	if err != nil {
		return nil, gerr.Internal("get task project failed, %s", err)
	}
	return ret, nil
}

func (t *TaskController) AddTasks(ctx context.Context, in *pb.TaskList) (*types.Empty, error) {
	return nil, nil
}

func (t *TaskController) RequestTasks(ctx context.Context, in *pb.RequestTaskOption) (*pb.TaskList, error) {
	return nil, nil
}

func (t *TaskController) CompleteTask(ctx context.Context, in *pb.TaskList) (*types.Empty, error) {
	return nil, nil
}

// projectKey redis Key
func (t *TaskController) projectKey(in *pb.TaskProject) string {
	return strings.Join([]string{rKeyPreTaskProject, in.Name}, rKeySep)
}
