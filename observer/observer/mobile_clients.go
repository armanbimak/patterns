package observer

import (
	"fmt"
	"github.com/armanbimak/patterns/observer/api"
)

func NewMobileClient(id string) api.IObserver {
	return onlineMobileClient{id: id}
}

type onlineMobileClient struct {
	id string
}

func (o onlineMobileClient) ID() string {
	return o.id
}

func (o onlineMobileClient) Update() {
	fmt.Println("updated MOBILE client")
}
