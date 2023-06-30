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
	ShowRobotStatus(data.RobotStatus)
	ShowError(string)
	ShowIdleStatus()
	StopChannel() chan struct{}
}

type Monitor interface {
	Start() error
	Stop()
	StatusChannel() chan data.RobotStatus
	ErrorChannel() chan error
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
	app := App{
		ui: ui.SysTray(ui.DefaultSystrayIcons{}),
		monitor: monitor.New(
			statussource.File(statussource.FileStatusSourceConfig{
				FileSource: filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
				Parser:     parser.JSON(),
				Rate:       1 * time.Second,
			}),
		),
	}

	go func() {
		for range app.ui.StopChannel() {
			app.monitor.Stop()
			app.ui.ShowIdleStatus()
		}
	}()

	go func() {
		for {
			select {
			case status := <-app.monitor.StatusChannel():
				app.ui.ShowRobotStatus(status)
			case err := <-app.monitor.ErrorChannel():
				app.ui.ShowError(err.Error())
			}
		}
	}()

	return app
}
