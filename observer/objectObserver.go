package observer

// type Observer interface {
// 	Observe(subject any)
// }

type Subject interface {
	Register(observer func(subject any))
	Deregister(observer func(subject any))
	Notify(subject any)
}
