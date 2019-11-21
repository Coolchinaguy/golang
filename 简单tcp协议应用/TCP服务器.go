package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, err2 := listener.Accept()
	if err != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//接受用户的请求

	buf := make([]byte, 1024) //1024的缓冲区大小
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("buf = ", string(buf[:n]))

	defer conn.Close() //关闭当前服务端链接
}
