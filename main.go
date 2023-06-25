package main

import (
	"robot-monitor/data"
	"robot-monitor/filesource"
	"robot-monitor/monitor"
	"robot-monitor/parser"
	"robot-monitor/ui"
	"time"

	"github.com/getlantern/systray"
)

type Monitor interface {
	Start() error
	Stop()
}

func main() {
	systrayUI := ui.SystrayUI{
		IdleIconPath:  "icons/robot-idle.ico",
		ErrorIconPath: "./icons/robot-error.ico",
		IconPathsByState: map[string]string{
			data.ROBOT_STATE_PASSING: "./icons/robot-passing.ico",
			data.ROBOT_STATE_FAILED:  "./icons/robot-failed.ico",
		},
	}

	idleMonitor := monitor.IdleMonitor{
		UI: &systrayUI,
	}

	activeMonitor := monitor.ActiveMonitor{
		Rate:       1 * time.Second,
		FileSource: filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
		Parser:     parser.JSON(),
		UI:         &systrayUI,
	}

	var currentMonitor Monitor = &activeMonitor

	systrayUI.OnStopClicked = func() {
		currentMonitor.Stop()
		currentMonitor = &idleMonitor
		currentMonitor.Start()
	}

	systray.Run(
		func() {
			systrayUI.Run()
			activeMonitor.Start()
		},
		func() {},
	)
}
