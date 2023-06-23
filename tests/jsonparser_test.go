package tests

import (
	"reflect"
	"robot-monitor/parser"
	"testing"
)

func TestJsonParser(t *testing.T) {
	jsonContent := `{
		"tasks": [
			{
				"name": "Task1",
				"status": "PASS"
			},
			{
				"name": "Task2",
				"status": "PASS"
			},
			{
				"name": "Task3",
				"status": "FAIL"
			},
			{
				"name": "Task4",
				"status": "SKIP"
			}
    	]
	}`

	expectedTasksStatus := map[string]uint16{
		"PASS": 2,
		"FAIL": 1,
		"SKIP": 1,
	}

	robotStatus, err := parser.JSON().Parse([]byte(jsonContent))
	if err != nil {
		t.Fatal(err)
	}

	tasksStatus := robotStatus.TasksStatus()

	if !reflect.DeepEqual(expectedTasksStatus, tasksStatus) {
		t.Fatalf("RobotStatus does not match expected!\nexpected: %v\nactual: %v", expectedTasksStatus, tasksStatus)
	}
}
