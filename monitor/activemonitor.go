package monitor

import (
	"robot-monitor/data"
	"time"
)

type FileSource interface {
	GetContent() ([]byte, error)
}

type Parser interface {
	Parse([]byte) (data.RobotStatus, error)
}

type ActiveUI interface {
	ShowRobotStatus(data.RobotStatus)
	ShowError(string)
}

type ActiveMonitor struct {
	Rate       time.Duration
	FileSource FileSource
	Parser     Parser
	UI         ActiveUI
	stop       chan struct{}
}

func (am *ActiveMonitor) Start() error {
	am.stop = make(chan struct{})
	ticker := time.NewTicker(am.Rate)

	// TODO: make filesource return a channel and let it implement watching
	go func() {
		for {
			select {
			case <-ticker.C:
				am.updateIcon()
			case <-am.stop:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}

func (am *ActiveMonitor) Stop() {
	am.stop <- struct{}{}
}

func (am *ActiveMonitor) updateIcon() {
	robotStatus, err := am.getRobotStatus()
	if err != nil {
		am.UI.ShowError(err.Error())
	} else {
		am.UI.ShowRobotStatus(robotStatus)
	}
}

func (am *ActiveMonitor) getRobotStatus() (robotStatus data.RobotStatus, err error) {
	content, err := am.FileSource.GetContent()
	if err != nil {
		return
	}

	return am.Parser.Parse(content)
}
