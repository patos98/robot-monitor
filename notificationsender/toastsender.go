package notificationsender

import (
	"robot-monitor/data"

	"github.com/go-toast/toast"
)

type ToastSender struct {
	appID string
}

func Toast(appID string) *ToastSender {
	return &ToastSender{
		appID: appID,
	}
}

func (t *ToastSender) Send(notification data.Notification) error {
	toastNotification := toast.Notification{
		AppID:   t.appID,
		Title:   notification.Title,
		Message: notification.Message,
	}

	return toastNotification.Push()
}
