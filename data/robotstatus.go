package data

import "fmt"

const TASK_STATUS_FAILED = "FAIL"

const ROBOT_STATE_PASSING = "PASSING"
const ROBOT_STATE_FAILED = "FAILED"

type Task struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RobotStatus struct {
	Tasks []Task `json:"tasks"`
}

func (rs *RobotStatus) TasksStatus() (tasksStatus map[string]uint16) {
	tasksStatus = make(map[string]uint16)
	for _, task := range rs.Tasks {
		tasksStatus[task.Status] += 1
	}
	return
}

func (rs *RobotStatus) String() string {
	return fmt.Sprint(rs.TasksStatus())
}

func (rs *RobotStatus) GetState() string {
	for _, task := range rs.Tasks {
		if task.Status == TASK_STATUS_FAILED {
			return ROBOT_STATE_FAILED
		}
	}
	return ROBOT_STATE_PASSING
}
