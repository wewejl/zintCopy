package net

import (
	"fmt"
	"net"
	"strings"
	"zinx/v1-basic-server/zinx/iface"
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
		fmt.Printf("net.ResolveTCPAddr err:",err)
		return
	}

	tcplistener,err:=net.ListenTCP(s.TCPVersion,tcpaddr)
	if err!=nil {
		fmt.Printf("net.ListenTCP err:",err)
		return
	}
	//2.建立链接Accept
	go func() {
		for  {
			tcpconn,err:=tcplistener.AcceptTCP()
			if err!=nil {
				fmt.Printf("net.AcceptTCP err:",err)
				return
			}
			fmt.Println("新链接成功")
			//3.对conn进行处理（业务）:接收client信息,转化为大写返回
			go func() {
				for  {
					//读取数据
					buf:=make([]byte,512)
					cnt,err:=tcpconn.Read(buf)
					if err!=nil {
						fmt.Printf("tcpconn.Read err",err)
						return
					}
					fmt.Println("Server <=== Content ",cnt,"data:",string(buf[:cnt]))
					//变成大写
					writBackInfo:=strings.ToUpper(string(buf[:cnt]))
					//服务器发送给客段
					cnt,err=tcpconn.Write([]byte(writBackInfo))
					fmt.Println("Server ===> Content",string(buf[:cnt]))
					//服务器发送给客户段
				}
			}()
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