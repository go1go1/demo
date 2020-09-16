package main

/**
Go语言的RPC规则：
 - 方法只能有两个可序列化的参数
 - 其中第二个参数是指针类型
 - 并且返回一个error类型
 - 同时必须是公开的方法
*/
type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}
