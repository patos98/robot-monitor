package main

import (
	"robot-monitor/data"
	"robot-monitor/filesource"
	"robot-monitor/monitor"
	"robot-monitor/parser"
	statussource "robot-monitor/statussource"
	"robot-monitor/ui"
	"time"
)

type UI interface {
	Run(func())
	ShowIdleStatus() error
	OnStopChannel() chan struct{}
}

type Monitor interface {
	Start() error
	Stop()
}

type App struct {
	ui      UI
	monitor Monitor
}

func main() {
	app := initializeApp()
	app.ui.Run(func() { app.monitor.Start() })
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

	statusSource := statussource.File(statussource.FileStatusSourceConfig{
		FileSource: filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
		Parser:     parser.JSON(),
		Rate:       1 * time.Second,
	})

	statusMonitor := monitor.New(statusSource, systrayUI)

	app := App{
		ui:      systrayUI,
		monitor: statusMonitor,
	}

	go func() {
		for range systrayUI.OnStopChannel() {
			app.monitor.Stop()
			app.ui.ShowIdleStatus()
		}
	}()

	return app
}
