package notifier

import "robot-monitor/data"

type Sender interface {
	Send(data.Notification) error
}
