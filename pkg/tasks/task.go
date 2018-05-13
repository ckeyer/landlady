package tasks

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/ckeyer/logrus"
	"github.com/funxdata/commons/gerr"
	pb "github.com/funxdata/landlady/proto"
	"github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	redis "gopkg.in/redis.v5"
)

const (
	rKeySep            = ":"
	rKeyPreTaskProject = "task_project"

	rKeyTodos = "todo"
	rKeyDoing = "doing"
	rKeyDone  = "done"
)

var _ pb.TasksServer = (*TaskController)(nil)

type TaskController struct {
	sync.Mutex
	*redis.Client
}

func NewTasksServer(rcli *redis.Client) *TaskController {
	return &TaskController{
		Client: rcli,
	}
}

func (t *TaskController) NewProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	in.Status = pb.TaskProjectStatus_Running
	in.StartAt = time.Now()

	bs, _ := json.Marshal(in)
	err := t.SetNX(t.projectKey(in.Name), bs, 0).Err()
	if err != nil {
		return nil, gerr.Internal("set task project %s failed, %s", in.Name, err)
	}

	return in, nil
}

func (t *TaskController) GetProject(ctx context.Context, in *pb.TaskProject) (*pb.TaskProject, error) {
	bs, err := t.Get(t.projectKey(in.Name)).Bytes()
	if err != nil {
		return nil, gerr.Internal("get task project failed, %s", err)
	}

	ret := &pb.TaskProject{}
	if err := json.Unmarshal(bs, &ret); err != nil {
		return nil, gerr.Internal("decode task project failed, %s", err)
	}
	return ret, nil
}

func (t *TaskController) PushTasks(ctx context.Context, in *pb.TaskList) (*types.Empty, error) {
	for _, item := range in.Items {
		if err := t.addTask(item); err != nil {
			logrus.Errorf("add task %s -> %s failed, %s", item.Url, item.ProjectName, err)
		}
	}
	return &types.Empty{}, nil
}

func (t *TaskController) CompletePushing(ctx context.Context, in *pb.TaskProject) (*types.Empty, error) {

	return &types.Empty{}, nil
}

func (t *TaskController) RequestTasks(ctx context.Context, in *pb.RequestTaskOption) (*pb.TaskList, error) {
	todoKey := t.projectKey(in.ProjectName, rKeyTodos)
	doingKey := t.projectKey(in.ProjectName, rKeyDoing)

	ret := make([]*pb.Task, 0, int(in.Count))
	err := t.Watch(func(ctx *redis.Tx) error {
		urls, _, err := ctx.SScan(todoKey, 0, "", in.Count).Result()
		if err != nil {
			return gerr.Internal("scan %s failed, %s", todoKey, err)
		}

		for _, url := range urls {
			ok, err := ctx.SMove(todoKey, doingKey, url).Result()
			if err != nil || !ok {
				logrus.Errorf("smove %q %s -> %s failed, %s", url, todoKey, doingKey, err)
				continue
			}

			ret = append(ret, &pb.Task{Url: url, ProjectName: in.ProjectName})
		}
		return nil
	}, todoKey, doingKey)
	if err != nil {
		return nil, err
	}

	return &pb.TaskList{Items: ret}, nil
}

func (t *TaskController) HandleTasks(ctx context.Context, in *pb.TaskList) (*types.Empty, error) {
	for _, task := range in.Items {
		doingKey := t.projectKey(task.ProjectName, rKeyDoing)
		doneKey := t.projectKey(task.ProjectName, rKeyDone)

		ok, err := t.SMove(doingKey, doneKey, task.Url).Result()
		if err != nil || !ok {
			logrus.Errorf("smove %q %s -> %s failed, %s", task.Url, doingKey, doneKey, err)
			continue
		}
	}
	return &types.Empty{}, nil
}

// addTask
func (t *TaskController) addTask(task *pb.Task) error {
	todoKey := t.projectKey(task.ProjectName, rKeyTodos)
	doingKey := t.projectKey(task.ProjectName, rKeyDoing)
	doneKey := t.projectKey(task.ProjectName, rKeyDone)

	return t.Watch(func(ctx *redis.Tx) error {
		if ctx.SIsMember(doingKey, task.Url).Val() ||
			ctx.SIsMember(doneKey, task.Url).Val() {
			return nil
		}
		return ctx.SAdd(todoKey, task.Url).Err()
	}, todoKey, doingKey, doneKey)
}

// projectKey redis Key: {key_prefix}:[modules:*]:{project_name}
func (t *TaskController) projectKey(prjName string, modules ...string) string {
	ks := make([]string, 0, len(modules)+2)
	ks = append([]string{rKeyPreTaskProject}, modules...)
	ks = append(ks, prjName)
	return strings.Join(ks, rKeySep)
}
