package monitor

import (
	"os"
	"robot-monitor/data"
	"time"

	"github.com/getlantern/systray"
)

const errorIconPath = "icons/robot-warning.ico"

var iconPathsByState = map[string]string{
	"PASSING": "icons/robot-passing.ico",
	"FAILED":  "icons/robot-failed.ico",
}

type FileSource interface {
	GetContent() ([]byte, error)
}

type Parser interface {
	Parse([]byte) (data.RobotStatus, error)
}

type ActiveMonitor struct {
	Rate       time.Duration
	FileSource FileSource
	Parser     Parser
	stop       <-chan struct{}
}

func (am *ActiveMonitor) Start() {
	am.stop = make(<-chan struct{})
	ticker := time.NewTicker(am.Rate)

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
}

func (am *ActiveMonitor) updateIcon() {
	robotStatus, err := am.getRobotStatus()
	if err != nil {
		icon, _ := os.ReadFile(errorIconPath)
		systray.SetIcon(icon)
		systray.SetTooltip(err.Error())
	} else {
		iconPath := iconPathsByState[robotStatus.GetState()]
		icon, _ := os.ReadFile(iconPath)
		systray.SetIcon(icon)
	}
}

func (am *ActiveMonitor) getRobotStatus() (robotStatus data.RobotStatus, err error) {
	content, err := am.FileSource.GetContent()
	if err != nil {
		return
	}

	return am.Parser.Parse(content)
}
