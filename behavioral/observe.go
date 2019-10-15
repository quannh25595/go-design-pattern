package behavioral

type (
	Event struct {
		Data int64
	}

	Subscriber interface {
		OnNotify(Event)
	}

	Publisher interface {
		Subscribe(Subscriber)
		Unsubscribe(Subscriber)
		Notify(Event)
	}
)
