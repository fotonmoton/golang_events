package catalog

import "fmt"

type Catalog struct {
	newArrivals []string
}

func NewCatalog() *Catalog {
	return &Catalog{}
}

func (c *Catalog) AddNewProduct(p string) {
	c.newArrivals = append(c.newArrivals, p)
}

func (c *Catalog) Observe(product any) {
	c.AddNewProduct(product.(string))

}

func (c *Catalog) Show() {
	if len(c.newArrivals) == 0 {
		fmt.Println("No new arrivals yet!")
		return
	}

	fmt.Println("here our new arrivals!")

	for _, p := range c.newArrivals {
		fmt.Println(p)
	}
}
