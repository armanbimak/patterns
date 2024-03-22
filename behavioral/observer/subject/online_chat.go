package subject

import (
	"github.com/armanbimak/patterns/behavioral/observer/api"
)

func NewOnlineChat() api.ISubject {
	return &onlineChat{
		observers: make(map[string]api.IObserver),
		messages:  make([]string, 0),
	}
}

type onlineChat struct {
	observers map[string]api.IObserver
	isChanged bool
	isMute    bool

	messages []string
}

func (o *onlineChat) Mute() {
	o.isMute = true
}

func (o *onlineChat) UnMute() {
	o.isMute = false
}

func (o *onlineChat) SendNewMessage(message string) {
	o.messages = append(o.messages, message)
	if !o.isMute {
		o.isChanged = true
	}

	o.NotifyObservers()
}

func (o *onlineChat) RegisterObserver(obs api.IObserver) {
	if _, ok := o.observers[obs.ID()]; !ok {
		o.observers[obs.ID()] = obs
	}
}

func (o *onlineChat) RemoveObserver(obs api.IObserver) {
	if _, ok := o.observers[obs.ID()]; ok {
		delete(o.observers, obs.ID())
	}
}

func (o *onlineChat) NotifyObservers() {
	if o.isChanged {
		for _, observer := range o.observers {
			observer.Update()
		}

		o.isChanged = false
	}
}
