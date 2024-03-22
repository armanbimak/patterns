package api

type ISubject interface {
	RegisterObserver(obs IObserver)
	RemoveObserver(obs IObserver)
	NotifyObservers()
}

type IObserver interface {
	ID() string
	Update()
}
