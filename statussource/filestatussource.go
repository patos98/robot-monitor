package statussource

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

type FileStatusSourceConfig struct {
	FileSource FileSource
	Parser     Parser
	Rate       time.Duration
}

type FileStatusSource struct {
	config        FileStatusSourceConfig
	statusChannel chan data.RobotStatus
	errorChannel  chan error
	stopChannel   chan struct{}
}

func File(config FileStatusSourceConfig) *FileStatusSource {
	statusSource := &FileStatusSource{
		config:        config,
		statusChannel: make(chan data.RobotStatus),
		errorChannel:  make(chan error),
		stopChannel:   make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-statusSource.stopChannel:
				return
			default:
				statusSource.loadRobotStatus()
				time.Sleep(config.Rate)
			}
		}
	}()

	return statusSource
}

func (robotStatusRource *FileStatusSource) StatusChannel() chan data.RobotStatus {
	return robotStatusRource.statusChannel
}

func (robotStatusRource *FileStatusSource) ErrorChannel() chan error {
	return robotStatusRource.errorChannel
}

func (robotStatusRource *FileStatusSource) Stop() {
	robotStatusRource.stopChannel <- struct{}{}
}

func (robotStatusRource *FileStatusSource) loadRobotStatus() {
	status, err := robotStatusRource.getRobotStatus()
	if err != nil {
		robotStatusRource.errorChannel <- err
	} else {
		robotStatusRource.statusChannel <- status
	}
}

func (robotStatusRource *FileStatusSource) getRobotStatus() (robotStatus data.RobotStatus, err error) {
	content, err := robotStatusRource.config.FileSource.GetContent()
	if err != nil {
		return
	}

	return robotStatusRource.config.Parser.Parse(content)
}
