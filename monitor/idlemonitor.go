package monitor

type IdleUI interface {
	ShowIdleStatus() error
}

type IdleMonitor struct {
	UI IdleUI
}

func (im *IdleMonitor) Start() (err error) {
	return im.UI.ShowIdleStatus()
}

func (im *IdleMonitor) Stop() {}
