package main

import (
	"events/catalog"
	"events/customers"
	"events/products"
	"events/pubsub"
	"fmt"
	"time"
)

func ObjectObserverExample() {
	greg := customers.NewCustomer("Greg", "blue jeans")
	bob := customers.NewCustomer("Bob", "t-shirt")
	newArrivals := catalog.NewCatalog()

	warehouse := products.NewWareHouse()

	warehouse.Register(greg)
	warehouse.Register(bob)
	warehouse.Register(newArrivals)

	newArrivals.Show()

	warehouse.AddProduct("t-shirt")
	warehouse.AddProduct("blue jeans")

	newArrivals.Show()
}

func PubSubExample() {
	topic := "productAdded"

	gregCustomer := customers.NewCustomer("Greg", "blue jeans")
	bobCustomer := customers.NewCustomer("Bob", "t-shirt")
	warehouse := products.NewWareHouse()
	newArrivals := catalog.NewCatalog()

	broker := pubsub.NewBroker()

	greg := broker.AddSubscriber()
	bob := broker.AddSubscriber()
	catalog := broker.AddSubscriber()

	broker.Subscribe(greg, topic)
	broker.Subscribe(bob, topic)
	broker.Subscribe(catalog, topic)

	newArrivals.Show()

	warehouse.AddProduct("t-shirt")
	broker.Publish(topic, "t-shirt")
	warehouse.AddProduct("blue jeans")
	broker.Publish(topic, "blue jeans")

	go func() {
		for message := range greg.Messages() {
			if gregCustomer.IsInterested(message.GetMessageBody()) {
				fmt.Println(gregCustomer.Name, "found what he interested in!: ", message.GetMessageBody())
			}
		}
	}()

	go func() {
		for message := range bob.Messages() {
			if bobCustomer.IsInterested(message.GetMessageBody()) {
				fmt.Println(bobCustomer.Name, "found what he interested in!: ", message.GetMessageBody())
			}
		}
	}()

	go func() {
		for message := range catalog.Messages() {
			newArrivals.AddNewProduct(message.GetMessageBody())
		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		newArrivals.Show()
	}()

	fmt.Scanln()
	fmt.Println("Done!")
}

func main() {
	ObjectObserverExample()
	PubSubExample()
}
