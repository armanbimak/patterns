package main

import (
	observer2 "github.com/armanbimak/patterns/behavioral/observer/observer"
	subject2 "github.com/armanbimak/patterns/behavioral/observer/subject"
)

func main() {
	chat := subject2.NewOnlineChat()

	desktopClient := observer2.NewDesktopClient("desktop")
	mobileClient := observer2.NewMobileClient("mobile")
	webClient := observer2.NewWebClient("web")

	chat.RegisterObserver(desktopClient)
	chat.RegisterObserver(mobileClient)
	chat.RegisterObserver(webClient)

	ch := chat.(subject2.IChat)

	chat.(subject2.IChat).SendNewMessage("message should be delivered")

	ch.Mute()

	chat.(subject2.IChat).SendNewMessage("message will not be delivered")

	ch.UnMute()

	chat.(subject2.IChat).SendNewMessage("message should be delivered")

}
