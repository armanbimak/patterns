package main

import (
	"github.com/armanbimak/patterns/observer/observer"
	"github.com/armanbimak/patterns/observer/subject"
)

func main() {
	chat := subject.NewOnlineChat()

	desktopClient := observer.NewDesktopClient("desktop")
	mobileClient := observer.NewMobileClient("mobile")
	webClient := observer.NewWebClient("web")

	chat.RegisterObserver(desktopClient)
	chat.RegisterObserver(mobileClient)
	chat.RegisterObserver(webClient)

	ch := chat.(subject.IChat)

	chat.(subject.IChat).SendNewMessage("message should be delivered")

	ch.Mute()

	chat.(subject.IChat).SendNewMessage("message will not be delivered")

	ch.UnMute()

	chat.(subject.IChat).SendNewMessage("message should be delivered")

}
