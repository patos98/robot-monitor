package monitor

import (
	"robot-monitor/data"
	"time"
)

type RobotStatusSource interface {
	StatusChannel() chan data.RobotStatus
	ErrorChannel() chan error
	Stop()
	Start(time.Duration)
}

type ActiveUI interface {
	ShowRobotStatus(data.RobotStatus)
	ShowError(string)
}

type ActiveMonitor struct {
	Rate              time.Duration
	RobotStatusRource RobotStatusSource
	UI                ActiveUI
}

func (activeMonitor *ActiveMonitor) Start() error {
	robotStatusSource := activeMonitor.RobotStatusRource
	robotStatusSource.Start(activeMonitor.Rate)
	go func() {
		for {
			select {
			case status := <-robotStatusSource.StatusChannel():
				activeMonitor.UI.ShowRobotStatus(status)
			case err := <-robotStatusSource.ErrorChannel():
				activeMonitor.UI.ShowError(err.Error())
				return
			}
		}
	}()

	return nil
}

func (activeMonitor *ActiveMonitor) Stop() {
	activeMonitor.RobotStatusRource.Stop()
}
