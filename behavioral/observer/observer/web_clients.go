package observer

import (
	"fmt"
	"github.com/armanbimak/patterns/behavioral/observer/api"
)

func NewWebClient(id string) api.IObserver {
	return onlineWebClient{id: id}
}

type onlineWebClient struct {
	id string
}

func (o onlineWebClient) ID() string {
	return o.id
}

func (o onlineWebClient) Update() {
	fmt.Println("updated WEB client")
}
