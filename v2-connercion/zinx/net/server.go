package net

import (
	"fmt"
	"net"
	"zinx/v2-connercion/zinx/iface"
)

//zinx是一个tcp的框架，使用socket进行服务搭建

//1.定义一个Server结构
type Server struct {
	//ip地址
	IP string
	//端口号
	Port uint32
	//服务器的名字
	Name string
	//tcpVersion
	TCPVersion string
}

//2.提供一个创建Server的方法
//func NewServer(name string) *Server  {
func NewServer(name string) iface.IServer  {
	return &Server{
		IP: "0.0.0.0",
		Port:8848,
		Name:name,
		TCPVersion:"tcp4", //tcp, tcp4,tcp6
	}

}

func (s *Server)Start()  {
	fmt.Println("[Zinx Server start....]")
	//TODO,标识以后在实现具体的逻辑


	//基于socker进行服务
	addr:=fmt.Sprintf("%s:%d",s.IP,s.Port)
	//1.创建socket 并进行监听
	tcpaddr,err:=net.ResolveTCPAddr(s.TCPVersion,addr)
	if err!=nil {
		fmt.Println("net.ResolveTCPAddr err:",err)
		return
	}

	tcplistener,err:=net.ListenTCP(s.TCPVersion,tcpaddr)
	if err!=nil {
		fmt.Println("net.ListenTCP err:",err)
		return
	}
	//创建connid,服务器分配，全局唯一，每次创建一个新的链接connid甲乙
	var connid uint32
	//2.建立链接Accept
	go func() {
		for  {
			tcpconn,err:=tcplistener.AcceptTCP()
			if err!=nil {
				fmt.Println("net.AcceptTCP err:",err,err)
				return
			}
			fmt.Println("新链接成功")

			//使用
			conn:=NewConnection(tcpconn,connid)
			//将所有连接相关的业务移动connection的Start函数中
			connid++
			go conn.Start()
			//3.对conn进行处理（业务）:接收client信息,转化为大写返回

		}
	}()
	//for  {
	//	;
	//}


}

func (s *Server)Stop()  {
	fmt.Println("[Zinx Server start....]")
	//TODO,标识以后在实现具体的逻辑
}

func (s *Server)Serve()  {
	fmt.Println("")
	s.Start()
	//阻塞，不占用cpu for循环是一直占用cpu
	select {}
	
}