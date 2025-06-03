package strategy

import "fmt"

type ChannelStrategy int

const (
	EmailStrategy ChannelStrategy = iota
	SMSStrategy
)

func (c ChannelStrategy) String() string {
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
	notifierMap map[ChannelStrategy]Notifier
}

func NewNotificationManager() *notificationManager {
	return &notificationManager{
		notifierMap: make(map[ChannelStrategy]Notifier),
	}
}

func (nm *notificationManager) RegisterNotifier(strategy ChannelStrategy, notifier Notifier) {
	nm.notifierMap[strategy] = notifier
}

func (nm *notificationManager) SendNotification(channel ChannelStrategy, to, message string) error {
	notifier, ok := nm.notifierMap[channel]
	if !ok {
		return fmt.Errorf("unsupported notification channel: %d", channel)
	}

	return notifier.Send(to, message)
}
