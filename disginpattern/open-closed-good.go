package main

import "fmt"

/*
* 开闭原则，对修改关闭，对扩展开放
 */

//抽象的银行业务员
type AbsBanker interface {
	DoBusi()
}
//抽象业务方法
func DoBusiness(banker AbsBanker) {
	banker.DoBusi()
}

//存款业务员
type SaveBanker struct {
	AbsBanker
}

func (b *SaveBanker) DoBusi() {
	fmt.Println("存款业务")
}

//转账业务员
type TransBanker struct {
	AbsBanker
}

func (b *TransBanker) DoBusi() {
	fmt.Println("转账业务")
}

//支付业务员
type PayBanker struct {
	AbsBanker
}

func (b *PayBanker) DoBusi() {
	fmt.Println("支付业务")
}

func main() {
	DoBusiness(&SaveBanker{})
	DoBusiness(&TransBanker{})
	DoBusiness(&PayBanker{})
}
