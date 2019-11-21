package main

import (
	"fmt"
	"net"
)

type Client struct {
	c    chan string //用户发送的数据
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户		cliAddr  =====>Client
var OnlineMap map[string]Client

var message = make(chan string)

//新开一个协程，转发消息，只要有消息来了，遍历变量map，给map给个成员都发送信息

func Manager() {
	//给map分配空间
	OnlineMap = make(map[string]Client)

	for {
		msg := <-message //没有信息时，通道阻塞

		//遍历map,当有信息时，给map每个成员都发送信息
		for _, cli := range OnlineMap {
			cli.c <- msg
		}
	}
}

func WriteMsgToclient(cli Client, conn net.Conn) {
	for msg := range cli.c { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg

	return
}

func HandleConn(conn net.Conn) { //处理用户链接
	defer conn.Close()
	//获取客户端的用户地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体,默认,用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	OnlineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToclient(cli, conn)

	//广播某个人在线

	message <- MakeMsg(cli, "login")

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历变量map，给map给个成员都发送信息

	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener err= ", err)
			continue
		}

		go HandleConn(conn) //处理用户连接
	}

}
