package strategy

import "fmt"

type ChanelStrategy int

const (
	EmailStrategy ChanelStrategy = iota
	SMSStrategy
)

func (c ChanelStrategy) String() string {
	switch c {
	case EmailStrategy:
		return "Email"
	case SMSStrategy:
		return "SMS"
	default:
		return "Unknown"
	}
}

type Notifier interface {
	Send(to, message string) error
}

type notificationManager struct {
	notifierMap map[ChanelStrategy]Notifier
}

func NewNotificationManager() (*notificationManager, error) {
	nm := &notificationManager{
		notifierMap: make(map[ChanelStrategy]Notifier),
	}

	return nm, nil
}

func (nm *notificationManager) RegisterNotifier(strategy ChanelStrategy, notifier Notifier) {
	nm.notifierMap[strategy] = notifier
}

func (nm *notificationManager) SendNotification(channel ChanelStrategy, to, message string) error {
	notifier, ok := nm.notifierMap[channel]
	if !ok {
		return fmt.Errorf("unsupported notification channel: %d", channel)
	}

	return notifier.Send(to, message)
}
