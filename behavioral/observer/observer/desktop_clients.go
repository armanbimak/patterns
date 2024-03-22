package observer

import (
	"fmt"
	"github.com/armanbimak/patterns/behavioral/observer/api"
)

func NewDesktopClient(id string) api.IObserver {
	return onlineDesktopClient{id: id}
}

type onlineDesktopClient struct {
	id string
}

func (o onlineDesktopClient) ID() string {
	return o.id
}

func (o onlineDesktopClient) Update() {
	fmt.Println("updated DESKTOP client")
}
