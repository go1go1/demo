package main

import "fmt"

// 抽象层 =============
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层 =============
type BenzCar struct {
}

func (c *BenzCar) Run() {
	fmt.Println("Benz run")
}

type BMWCar struct {
}

func (c *BMWCar) Run() {
	fmt.Println("BWM run")
}

type ToyotaCar struct {
}

func (c *ToyotaCar) Run() {
	fmt.Println("Toyota run")
}

type DZhang struct {
}

func (d *DZhang) Drive(car Car) {
	fmt.Println("DZhang drive: ")
	car.Run()
}

type DLi struct {
}

func (d *DLi) Drive(car Car) {
	fmt.Println("DLi drive: ")
	car.Run()
}

// 业务逻辑层 =============
func main() {
	var c Car
	c = &BenzCar{}

	var d Driver
	d = &DZhang{}

	d.Drive(c)
}
