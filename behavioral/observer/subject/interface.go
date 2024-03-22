package subject

type IChat interface {
	SendNewMessage(message string)
	Mute()
	UnMute()
}
