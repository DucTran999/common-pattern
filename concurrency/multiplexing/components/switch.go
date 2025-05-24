package components

import (
	"log"
)

type UnicastChan chan string
type BroadcastChan = UnicastChan

type Switch interface {
	BroadcastChan() BroadcastChan
	RegisterDeviceUnicast(unicastChannels ...UnicastChan)
	Listen()
}

type switchDevice struct {
	broadcast       BroadcastChan
	switchUnicast   UnicastChan
	routerUnicast   UnicastChan
	unicastChannels []UnicastChan
	arpCache        map[string]string
}

func NewSwitch(
	broadcast BroadcastChan, ackChan UnicastChan, routerChan UnicastChan,
) *switchDevice {
	return &switchDevice{
		broadcast:     broadcast,
		switchUnicast: ackChan,
		routerUnicast: routerChan,
		arpCache:      make(map[string]string),
	}
}

func (s *switchDevice) BroadcastChan() BroadcastChan {
	return s.broadcast
}

func (s *switchDevice) RegisterDeviceUnicast(unicastChannels ...UnicastChan) {
	s.unicastChannels = append(s.unicastChannels, unicastChannels...)
}

func (s *switchDevice) Listen() {
	for askedIP := range s.broadcast {
		// each arp request send to all devices in network to ask for MAC address
		log.Printf("Switch: Who has IP %s?", askedIP)
		for _, u := range s.unicastChannels {
			u <- askedIP
		}

		// Listen ack from devices and reply to router
		deviceAckCount := 0
		for ack := range s.switchUnicast {
			deviceAckCount++
			if ack != "" {
				s.arpCache[askedIP] = ack
				log.Printf("Switch: cached IP: %s - MAC: %s", askedIP, ack)
			}

			if len(s.unicastChannels) == deviceAckCount {
				s.routerUnicast <- s.arpCache[askedIP]
				log.Printf("Switch: received all ack")
				break
			}
		}
	}

	for _, u := range s.unicastChannels {
		close(u)
	}
	close(s.switchUnicast)
}
