package main

import (
	"robot-monitor/data"
	"robot-monitor/filesource"
	"robot-monitor/monitor"
	"robot-monitor/notificationsender"
	"robot-monitor/notifier"
	"robot-monitor/parser"
	"robot-monitor/statussource"
	"robot-monitor/ui"
	"time"
)

type UI interface {
	Run(func())
	ShowRobotStatus(data.RobotStatus)
	ShowError(string)
	ShowIdleStatus()
	StartChannel() chan struct{}
	StopChannel() chan struct{}
}

type Monitor interface {
	Start() error
	Stop()
	StatusChannel() chan data.RobotStatus
	ErrorChannel() chan error
}

type Notifier interface {
	ShouldNotify(data.RobotStatus) bool
	Notify(notifier.Sender) error
}

type NotificationConfig struct {
	notifiers           []Notifier
	notificationSenders []notifier.Sender
}

type App struct {
	ui                  UI
	monitor             Monitor
	notificationConfigs []NotificationConfig
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
		notificationConfigs: []NotificationConfig{{
			notifiers:           []Notifier{notifier.FirstFailed()},
			notificationSenders: []notifier.Sender{notificationsender.Toast("Robot monitor")},
		}},
	}

	go func() {
		for range app.ui.StopChannel() {
			app.monitor.Stop()
			app.ui.ShowIdleStatus()
		}
	}()

	go func() {
		for range app.ui.StartChannel() {
			app.monitor.Start()
		}
	}()

	go func() {
		for {
			select {
			case status := <-app.monitor.StatusChannel():
				app.ui.ShowRobotStatus(status)
				for _, notificationConfig := range app.notificationConfigs {
					for _, notifier := range notificationConfig.notifiers {
						if notifier.ShouldNotify(status) {
							for _, sender := range notificationConfig.notificationSenders {
								notifier.Notify(sender)
							}
						}
					}
				}
			case err := <-app.monitor.ErrorChannel():
				app.ui.ShowError(err.Error())
			}
		}
	}()

	return app
}
