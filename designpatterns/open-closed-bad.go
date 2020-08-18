package main

import "fmt"

type Banker struct {
}

type BankBusi interface {
	save()
	trans()
	pay()
}

func (b *Banker) save() {
	fmt.Println("存款")
}

func (b *Banker) trans() {
	fmt.Println("转账")
}

func (b *Banker) pay() {
	fmt.Println("支付")
}

func main() {
	b := &Banker{}

	b.save()
	b.trans()
	b.pay()
}
