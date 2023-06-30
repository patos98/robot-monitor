package ui

import "github.com/getlantern/systray"

type MenuItemConfig struct {
	title   string
	tooltip string
	onClick func()
}

func (systrayUI *SystrayUI) initMenu() {
	systrayUI.startMenuItem = systrayUI.addMenuItem(MenuItemConfig{
		title:   "Start",
		tooltip: "Start monitoring robot file",
		onClick: func() { systrayUI.startChannel <- struct{}{} },
	})

	systrayUI.stopMenuItem = systrayUI.addMenuItem(MenuItemConfig{
		title:   "Stop",
		tooltip: "Stop monitoring robot file",
		onClick: func() { systrayUI.stopChannel <- struct{}{} },
	})

	systrayUI.addMenuItem(MenuItemConfig{
		title:   "Quit",
		tooltip: "Quit the whole app",
		onClick: func() { systray.Quit() },
	})
}

func (systrayUI *SystrayUI) addMenuItem(menuItemConfig MenuItemConfig) *systray.MenuItem {
	menuItem := systray.AddMenuItem(menuItemConfig.title, menuItemConfig.tooltip)

	go func() {
		for {
			<-menuItem.ClickedCh
			menuItemConfig.onClick()
		}
	}()

	return menuItem
}
