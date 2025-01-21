package main

import "fmt"

// type Driven interface {
// 	Drive()
// }

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("drive")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
		return
	}
	fmt.Println("driver too young")
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		Car{},
		driver,
	}
}

func main() {
	car1 := NewCarProxy(&Driver{15})
	car2 := NewCarProxy(&Driver{17})

	car1.Drive()
	car2.Drive()

}
