package strategy

type Notifier interface {
	SendNotification(to, message string) error
}

type notificationManager struct {
	NotifierMap map[string]Notifier
}

func NewNotificationManager() (*notificationManager, error) {
	nm := &notificationManager{
		NotifierMap: make(map[string]Notifier, 0),
	}

	return nm, nil
}

func (nm *notificationManager) RegisterNotifier(strategy string, notifier Notifier) {
	nm.NotifierMap[strategy] = notifier
}
