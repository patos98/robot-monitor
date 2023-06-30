package notifier

import (
	"github.com/go-toast/toast"
)

type Notification interface {
	Title() string
	Message() string
}

type ToastNotifier struct {
	appID    string
	iconPath string
}

func Toast(appID string, iconPath string) *ToastNotifier {
	return &ToastNotifier{
		appID:    appID,
		iconPath: iconPath,
	}
}

func (t *ToastNotifier) Notify(notification Notification) error {
	toastNotification := toast.Notification{
		AppID:   t.appID,
		Icon:    t.iconPath,
		Title:   notification.Title(),
		Message: notification.Message(),
	}

	return toastNotification.Push()
}
