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
	ShowIdleStatus() error
	OnStopChannel() chan struct{}
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
		ui: ui.SysTray(ui.SystrayUIConfig{
			IdleIconPath:  "icons/robot-idle.ico",
			ErrorIconPath: "./icons/robot-warning.ico",
			IconPathsByState: map[string]string{
				data.ROBOT_STATE_PASSING: "./icons/robot-passing.ico",
				data.ROBOT_STATE_FAILED:  "./icons/robot-failed.ico",
			},
		}),
		monitor: monitor.New(
			statussource.File(statussource.FileStatusSourceConfig{
				FileSource: filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
				Parser:     parser.JSON(),
				Rate:       1 * time.Second,
			}),
		),
	}

	go func() {
		for range app.ui.OnStopChannel() {
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
