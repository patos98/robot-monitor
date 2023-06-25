package ui

import (
	"os"
	"robot-monitor/data"

	"github.com/getlantern/systray"
)

type SystrayUI struct {
	IdleIconPath     string
	ErrorIconPath    string
	IconPathsByState map[string]string
	OnStopClicked    func()
	stopMenuItem     *systray.MenuItem
	quitMenuItem     *systray.MenuItem
}

func (s *SystrayUI) Run() {
	s.stopMenuItem = systray.AddMenuItem("Stop", "Stop monitoring robot file")
	go func() {
		for {
			<-s.stopMenuItem.ClickedCh
			s.OnStopClicked()
		}
	}()
	s.quitMenuItem = systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-s.quitMenuItem.ClickedCh
		systray.Quit()
	}()
}

func (s *SystrayUI) ShowIdleStatus() error {
	s.stopMenuItem.Hide()
	return updateSysTrayIcon(s.IdleIconPath, "Click to set robot file to monitor.")
}

func (s *SystrayUI) ShowRobotStatus(robotStatus data.RobotStatus) {
	iconPath := s.IconPathsByState[robotStatus.GetState()]
	updateSysTrayIcon(iconPath, robotStatus.String())
}

func (s *SystrayUI) ShowError(errorMessage string) {
	s.stopMenuItem.Show()
	updateSysTrayIcon(s.ErrorIconPath, errorMessage) // TODO: handle error
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
