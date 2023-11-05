package products

import (
	"events/observer"
)

type Warehouse struct {
	products    []string
	subscribers []observer.Observer
}

func NewWareHouse() *Warehouse {
	return &Warehouse{}
}

func (w *Warehouse) AddProduct(product string) {
	w.products = append(w.products, product)
	w.Notify(product)
}

func (w *Warehouse) Register(listener observer.Observer) {
	w.subscribers = append(w.subscribers, listener)
}

// Not sure it works
func (w *Warehouse) Deregister(listener observer.Observer) {
	// idx := slices.IndexFunc(w.subscribers, func(o observer.Observer) bool { return o == listener })
	// w.subscribers = slices.Delete(w.subscribers, idx, idx)
}

func (w *Warehouse) Notify(subject any) {
	for _, listener := range w.subscribers {
		listener(subject)
	}
}
