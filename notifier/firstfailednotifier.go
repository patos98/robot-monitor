package notifier

import "robot-monitor/data"

type FirstFailedNotifier struct {
	lastStatus data.RobotStatus
}

func FirstFailed() *FirstFailedNotifier {
	return &FirstFailedNotifier{}
}

func (failedNotifier *FirstFailedNotifier) ShouldNotify(status data.RobotStatus) bool {
	lastState := failedNotifier.lastStatus.GetState()
	failedNotifier.lastStatus = status
	return lastState != data.ROBOT_STATE_FAILED && status.GetState() == data.ROBOT_STATE_FAILED
}

func (failedNotifier *FirstFailedNotifier) Notify(sender Sender) error {
	return sender.Send(data.Notification{
		Title:   "Tests failed",
		Message: "Failed task: " + failedNotifier.lastStatus.GetFirstFailedTask(),
	})
}
