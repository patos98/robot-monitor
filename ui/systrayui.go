package ui

import (
	"robot-monitor/data"

	"github.com/getlantern/systray"
)

type SystrayUI struct {
	icons         SystrayIcons
	onStopChannel chan struct{}
	stopMenuItem  *systray.MenuItem
}

func SysTray(icons SystrayIcons) *SystrayUI {
	return &SystrayUI{
		icons:         icons,
		onStopChannel: make(chan struct{}),
	}
}

func (systrayUI *SystrayUI) Run(onReady func()) {
	systray.Run(
		func() {
			systrayUI.initMenu()
			onReady()
		},
		func() {},
	)
}

func (systrayUI *SystrayUI) ShowIdleStatus() {
	systrayUI.stopMenuItem.Hide()
	updateSysTrayIcon(systrayUI.icons.Idle(), "Click to set robot file to monitor.")
}

func (systrayUI *SystrayUI) ShowRobotStatus(robotStatus data.RobotStatus) {
	icon, err := systrayUI.icons.ForStatus(robotStatus)
	if err != nil {
		systrayUI.ShowError(err.Error())
	} else {
		updateSysTrayIcon(icon, robotStatus.String())
	}
}

func (systrayUI *SystrayUI) ShowError(errorMessage string) {
	systrayUI.stopMenuItem.Show()
	updateSysTrayIcon(systrayUI.icons.Error(), errorMessage)
}

func (s *SystrayUI) StopChannel() chan struct{} {
	return s.onStopChannel
}

func updateSysTrayIcon(icon []byte, tooltip string) {
	systray.SetIcon(icon)
	systray.SetTooltip(tooltip)
}
