package task

var globalTaskMap map[string]Task

type Task interface {
	Name() string
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