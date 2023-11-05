package observer

type Observer func(subject any)

type CallbackSubject interface {
	Register(observer func(subject any))
	Deregister(observer int)
	Notify(subject any)
}
