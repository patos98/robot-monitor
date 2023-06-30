package monitor

import (
	"robot-monitor/data"
)

type StatusSource interface {
	StatusChannel() chan data.RobotStatus
	ErrorChannel() chan error
}

type UI interface {
	ShowRobotStatus(data.RobotStatus)
	ShowError(string)
}

type Monitor struct {
	statusSource StatusSource
	ui           UI
	stopChannel  chan struct{}
}

func New(statusSource StatusSource, ui UI) *Monitor {
	return &Monitor{
		statusSource: statusSource,
		ui:           ui,
		stopChannel:  make(chan struct{}),
	}
}

func (monitor *Monitor) Start() error {
	statusSource := monitor.statusSource

	go func() {
		for {
			select {
			case status := <-statusSource.StatusChannel():
				monitor.ui.ShowRobotStatus(status)
			case err := <-statusSource.ErrorChannel():
				monitor.ui.ShowError(err.Error())
			case <-monitor.stopChannel:
				return
			}
		}
	}()

	return nil
}

func (monitor *Monitor) Stop() {
	monitor.stopChannel <- struct{}{}
}
