package main

//使用接口实现多态
import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (n *NokiaPhone) call() {
	fmt.Println("nokia phone call...")
}

type ApplePhone struct {
}

func (a *ApplePhone) call() {
	fmt.Println("apple phone call...")
}

func callPhone(phone Phone) {
	phone.call()
}

func main() {
	callPhone(&NokiaPhone{})
	callPhone(&ApplePhone{})
}
