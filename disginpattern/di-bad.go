package main

import "fmt"

type Benz struct {
}

func (c *Benz) Run() {
	fmt.Println("Benz run")
}

type BMW struct {
}

func (c *BMW) Run() {
	fmt.Println("BMW run")
}

type DriverZhang struct {
}

func (d *DriverZhang) driveBenz(benz *Benz) {
	fmt.Println("DriverZhang dirve")
	benz.Run()
}

func (d *DriverZhang) driveBMW(bmw *BMW) {
	fmt.Println("DriverZhang dirve")
	bmw.Run()
}

type DriverLi struct {
}

func (d *DriverLi) driveBenz(benz *Benz) {
	fmt.Println("DriverZhang dirve")
	benz.Run()
}
func (d *DriverLi) driveBMW(bmw *BMW) {
	fmt.Println("DriverZhang dirve")
	bmw.Run()
}

func main() {
	d1 := &DriverZhang{}
	d1.driveBenz(&Benz{})
	d1.driveBMW(&BMW{})

	d2 := &DriverLi{}
	d2.driveBenz(&Benz{})
	d2.driveBMW(&BMW{})
}
