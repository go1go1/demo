package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("connection failed.", err)
	}

	defer socket.Close()

	//发送数据
	sendData := []byte("hello server")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("send data failed.", err)
		return
	}

	//接收数据
	data := make([]byte, 4096)
	count, addr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data failed,", err)
		return
	}
	fmt.Printf("read data: %s, addr:%v, count:%d\n\n", string(data[0:count]), addr, count)
}
