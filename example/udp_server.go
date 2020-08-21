package main

import (
	"fmt"
	"net"
)

func main() {
	//创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listening failed.", err)
		return
	}

	fmt.Println("listening started...")

	defer socket.Close()

	for {
		//读取数据
		data := make([]byte, 4096)
		count, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read data failed", err)
			continue
		}
		fmt.Printf("read data: %s, addr:%v, count:%d\n\n", string(data[0:count]), addr, count)

		//发送数据
		sendData := []byte("hello client")
		_, err = socket.WriteToUDP(sendData, addr)
		if err != nil {
			fmt.Println("send data failed, ", err)
			return
		}
	}
}
