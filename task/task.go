package task

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var globalTaskMap map[string]Task
var TaskConfDir string
var TaskLogDir string

type Task interface {
	Name() string
	Enabled() bool
	Load() error
	Run() error
	Stop()
}

func RegisterTask(task Task) {
	if globalTaskMap == nil {
		globalTaskMap = make(map[string]Task)
	}
	globalTaskMap[task.Name()] = task
}

func LoadAllTasks() map[string]Task {
	return globalTaskMap
}

func ReadTaskConfig(t Task) error {
	path := fmt.Sprintf("%s/%s.toml",
		TaskConfDir, t.Name())
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = toml.Decode(string(data), t)
	return err
}
