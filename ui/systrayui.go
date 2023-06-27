package ui

import (
	"os"
	"robot-monitor/data"

	"github.com/getlantern/systray"
)

type SystrayUIConfig struct {
	IdleIconPath     string
	ErrorIconPath    string
	IconPathsByState map[string]string
}

type SystrayUI struct {
	config        SystrayUIConfig
	onStopChannel chan struct{}
	stopMenuItem  *systray.MenuItem
	quitMenuItem  *systray.MenuItem
}

func SysTray(config SystrayUIConfig) SystrayUI {
	return SystrayUI{
		config:        config,
		onStopChannel: make(chan struct{}),
	}
}

func (s *SystrayUI) Run() {
	systray.Run(
		func() {
			s.stopMenuItem = systray.AddMenuItem("Stop", "Stop monitoring robot file")
			go func() {
				for {
					<-s.stopMenuItem.ClickedCh
					s.onStopChannel <- struct{}{}
				}
			}()
			s.quitMenuItem = systray.AddMenuItem("Quit", "Quit the whole app")
			go func() {
				<-s.quitMenuItem.ClickedCh
				systray.Quit()
			}()
		},
		func() {},
	)
}

func (s *SystrayUI) ShowIdleStatus() error {
	s.stopMenuItem.Hide()
	return updateSysTrayIcon(s.config.IdleIconPath, "Click to set robot file to monitor.")
}

func (s *SystrayUI) ShowRobotStatus(robotStatus data.RobotStatus) {
	iconPath := s.config.IconPathsByState[robotStatus.GetState()]
	updateSysTrayIcon(iconPath, robotStatus.String())
}

func (s *SystrayUI) ShowError(errorMessage string) {
	s.stopMenuItem.Show()
	updateSysTrayIcon(s.config.ErrorIconPath, errorMessage) // TODO: handle error
}

func (s *SystrayUI) OnStopChannel() chan struct{} {
	return s.onStopChannel
}

func updateSysTrayIcon(iconPath string, tooltip string) (err error) {
	icon, err := os.ReadFile(iconPath) // TODO: handle error
	if err != nil {
		return
	}
	systray.SetIcon(icon)
	systray.SetTooltip(tooltip)
	return
}
