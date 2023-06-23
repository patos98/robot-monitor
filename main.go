package main

import (
	"robot-monitor/monitor"

	"github.com/getlantern/systray"
)

func main() {
	idleMonitor := monitor.Idle()
	currentMonitor := idleMonitor

	systray.Run(
		func() {
			currentMonitor.Start()
			mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
			go func() {
				<-mQuitOrig.ClickedCh
				systray.Quit()
			}()
		},
		func() {},
	)
}
