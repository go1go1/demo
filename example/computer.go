package main

import "fmt"

//抽象层
type Cpu interface {
	Caculate()
}

type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type Computer struct {
	cpu    Cpu
	card   Card
	memory Memory
}

//实现层

type InterCpu struct {
	Cpu
}

func (i *InterCpu) Caculate() {
	fmt.Println("Inter CPU...")
}

type InterCard struct {
	Card
}

func (i *InterCard) Display() {
	fmt.Println("Inter Card")
}

type InterMemory struct {
	Memory
}

func (i *InterMemory) Storage() {
	fmt.Println("Inter Memory")
}

func NewComputer(cpu Cpu, card Card, memory Memory) *Computer {
	return &Computer{
		cpu:    cpu,
		memory: memory,
		card:   card,
	}
}

func (c *Computer) Work() {
	c.cpu.Caculate()
	c.memory.Storage()
	c.card.Display()
}

func main() {
	com1 := NewComputer(&InterCpu{}, &InterCard{}, &InterMemory{})

	com1.Work()
}
