package main

import "zinx/v1-basic-server/zinx/net"

//里面实现对zinx框架的使用，创建一个自己的服务器

func main()  {
	server:=net.NewServer("zinxV1.0")
	server.Serve()
}
