package ui

import (
	"robot-monitor/data"

	"github.com/getlantern/systray"
)

type SystrayUI struct {
	icons         SystrayIcons
	startChannel  chan struct{}
	stopChannel   chan struct{}
	startMenuItem *systray.MenuItem
	stopMenuItem  *systray.MenuItem
}

func SysTray(icons SystrayIcons) *SystrayUI {
	return &SystrayUI{
		icons:        icons,
		startChannel: make(chan struct{}),
		stopChannel:  make(chan struct{}),
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
	systrayUI.startMenuItem.Show()
	updateSysTrayIcon(systrayUI.icons.Idle(), "Click to set robot file to monitor.")
}

func (systrayUI *SystrayUI) ShowRobotStatus(robotStatus data.RobotStatus) {
	systrayUI.stopMenuItem.Show()
	systrayUI.startMenuItem.Hide()
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

func (s *SystrayUI) StartChannel() chan struct{} {
	return s.startChannel
}

func (s *SystrayUI) StopChannel() chan struct{} {
	return s.stopChannel
}

func updateSysTrayIcon(icon []byte, tooltip string) {
	systray.SetIcon(icon)
	systray.SetTooltip(tooltip)
}
