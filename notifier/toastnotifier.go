package notifier

import (
	"github.com/go-toast/toast"
)

type Notification struct {
	Title   string
	Message string
}

type Toast struct {
	AppID    string
	IconPath string
}

func (t *Toast) Notify(notification Notification) error {
	toastNotification := toast.Notification{
		AppID:   t.AppID,
		Icon:    t.IconPath,
		Title:   notification.Title,
		Message: notification.Message,
	}

	return toastNotification.Push()
}
