package main

import (
	"robot-monitor/filesource"
	"robot-monitor/monitor"
	"robot-monitor/parser"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	// idleMonitor := monitor.IdleMonitor{}
	activeMonitor := monitor.ActiveMonitor{
		Rate:       1 * time.Second,
		FileSource: filesource.Local("F:/Users/szpat/Downloads/robot-monitor/tasks-status.json"),
		Parser:     parser.JSON(),
	}
	// currentMonitor := idleMonitor

	systray.Run(
		func() {
			activeMonitor.Start()
			mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
			go func() {
				<-mQuitOrig.ClickedCh
				systray.Quit()
			}()
		},
		func() {},
	)
}
