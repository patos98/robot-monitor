package ui

import (
	_ "embed"
	"errors"
	"robot-monitor/data"
)

type SystrayIcons interface {
	Idle() []byte
	Error() []byte
	ForStatus(data.RobotStatus) ([]byte, error)
}

type DefaultSystrayIcons struct{}

//go:embed icons/robot-idle.ico
var idleIcon []byte

//go:embed icons/robot-warning.ico
var errorIcon []byte

//go:embed icons/robot-passing.ico
var passingIcon []byte

//go:embed icons/robot-failed.ico
var failedIcon []byte

func (DefaultSystrayIcons) Idle() []byte {
	return idleIcon
}

func (DefaultSystrayIcons) Error() []byte {
	return errorIcon
}

func (DefaultSystrayIcons) ForStatus(status data.RobotStatus) (icon []byte, err error) {
	switch status.GetState() {
	case data.ROBOT_STATE_PASSING:
		icon = passingIcon
	case data.ROBOT_STATE_FAILED:
		icon = failedIcon
	default:
		err = errors.New("Icon not set for status: " + status.GetState())
	}

	return
}
