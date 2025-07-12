package task

type TaskManager struct {
	tasks  map[int]Task
	lastID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[int]Task),
	}
}

func (tm *TaskManager) Add(description string) (int, error) {
	task, err := NewTask(tm.lastID+1, description)

	if err != nil {
		return 0, err
	}

	tm.lastID++
	tm.tasks[task.ID] = *task
	return task.ID, nil
}
