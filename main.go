package main

import (
	"robot-monitor/data"
	"robot-monitor/filesource"
	"robot-monitor/monitor"
	"robot-monitor/parser"
	"robot-monitor/robotstatussource"
	"robot-monitor/ui"
	"time"
)

type UI interface {
	Run()
	OnStopChannel() chan struct{}
}

type Monitor interface {
	Start() error
	Stop()
}

type App struct {
	ui             UI
	currentMonitor Monitor
}

func main() {
	app := initializeApp()
	app.currentMonitor.Start()
	app.ui.Run()
}

func initializeApp() App {
	systrayUI := ui.SysTray(ui.SystrayUIConfig{
		IdleIconPath:  "icons/robot-idle.ico",
		ErrorIconPath: "./icons/robot-error.ico",
		IconPathsByState: map[string]string{
			data.ROBOT_STATE_PASSING: "./icons/robot-passing.ico",
			data.ROBOT_STATE_FAILED:  "./icons/robot-failed.ico",
		},
	})

	idleMonitor := monitor.IdleMonitor{UI: &systrayUI}

	robotStatusSource := robotstatussource.File(
		filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
		parser.JSON(),
	)

	activeMonitor := monitor.ActiveMonitor{
		Rate:              1 * time.Second,
		RobotStatusRource: &robotStatusSource,
		UI:                &systrayUI,
	}

	app := App{
		ui:             &systrayUI,
		currentMonitor: &activeMonitor,
	}

	go func() {
		for range systrayUI.OnStopChannel() {
			app.currentMonitor.Stop()
			app.currentMonitor = &idleMonitor
			app.currentMonitor.Start()
		}
	}()

	return app
}
