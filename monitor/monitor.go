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
	statusSource  StatusSource
	statusChannel chan data.RobotStatus
	errorChannel  chan error
	stopChannel   chan struct{}
}

func New(statusSource StatusSource) *Monitor {
	return &Monitor{
		statusSource:  statusSource,
		statusChannel: make(chan data.RobotStatus),
		errorChannel:  make(chan error),
		stopChannel:   make(chan struct{}),
	}
}

func (monitor *Monitor) Start() error {
	statusSource := monitor.statusSource

	go func() {
		for {
			select {
			case status := <-statusSource.StatusChannel():
				monitor.statusChannel <- status
			case err := <-statusSource.ErrorChannel():
				monitor.errorChannel <- err
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

func (monitor *Monitor) StatusChannel() chan data.RobotStatus {
	return monitor.statusChannel
}

func (monitor *Monitor) ErrorChannel() chan error {
	return monitor.errorChannel
}
