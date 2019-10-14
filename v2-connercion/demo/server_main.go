package main

import "zinx/v2-connercion/zinx/net"

//里面实现对zinx框架的使用，创建一个自己的服务器

func main()  {
	server:=net.NewServer("zinxV1.0")
	server.Serve()
}
