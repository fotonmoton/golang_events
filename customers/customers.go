package customers

import (
	"fmt"
	"slices"
)

type Customer struct {
	Name      string
	interests []string
}

func NewCustomer(name string, interests ...string) *Customer {
	return &Customer{
		Name: name, interests: interests,
	}
}

func (c *Customer) Observe(subject any) {
	if c.IsInterested(subject.(string)) {
		fmt.Println(c.Name, "found what he interested in!: ", subject)
	}
}

func (c *Customer) IsInterested(subject string) bool {
	return slices.Contains(c.interests, subject)
}
