package robotstatussource

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

type FileRobotStatusSource struct {
	fileSource    FileSource
	parser        Parser
	statusChannel chan data.RobotStatus
	errorChannel  chan error
	stopChannel   chan struct{}
}

func File(fileSource FileSource, parser Parser) FileRobotStatusSource {
	return FileRobotStatusSource{
		fileSource:    fileSource,
		parser:        parser,
		statusChannel: make(chan data.RobotStatus),
		errorChannel:  make(chan error),
		stopChannel:   make(chan struct{}),
	}
}

func (robotStatusRource *FileRobotStatusSource) StatusChannel() chan data.RobotStatus {
	return robotStatusRource.statusChannel
}

func (robotStatusRource *FileRobotStatusSource) ErrorChannel() chan error {
	return robotStatusRource.errorChannel
}

func (robotStatusRource *FileRobotStatusSource) Stop() {
	robotStatusRource.stopChannel <- struct{}{}
}

func (robotStatusRource *FileRobotStatusSource) Start(rate time.Duration) {
	ticker := time.NewTicker(rate)
	go func() {
		for {
			select {
			case <-ticker.C:
				status, err := robotStatusRource.getRobotStatus()
				if err != nil {
					robotStatusRource.errorChannel <- err
				} else {
					robotStatusRource.statusChannel <- status
				}
			case <-robotStatusRource.stopChannel:
				ticker.Stop()
				return
			}
		}
	}()
}

func (robotStatusRource *FileRobotStatusSource) getRobotStatus() (robotStatus data.RobotStatus, err error) {
	content, err := robotStatusRource.fileSource.GetContent()
	if err != nil {
		return
	}

	return robotStatusRource.parser.Parse(content)
}
