package monitor

import (
	"os"

	"github.com/getlantern/systray"
)

type IdleMonitor struct{}

func Idle() *IdleMonitor {
	return &IdleMonitor{}
}

func (im *IdleMonitor) Start() (err error) {
	icon, err := os.ReadFile("icons/robot-idle.ico")
	if err != nil {
		return
	}
	systray.SetTooltip("Click to set robot file to monitor.")
	systray.SetIcon(icon)
	return
}

func (im *IdleMonitor) Stop() {}
