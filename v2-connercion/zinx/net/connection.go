package net

import (
	"fmt"
	"net"
	"strings"
)

type Connection struct {
	
	tcpconn *net.TCPConn
	
	//2.每一个链接都有一个唯一的id
	connid uint32
} 

//提供创建Connection的方法
func NewConnection(conn *net.TCPConn,cnnid uint32) *Connection {
	return &Connection{
		tcpconn: conn,
		connid:  cnnid,
	}
}

//返回原生tcpconn
func (c *Connection)GetTcpConn() *net.TCPConn  {
	return c.tcpconn
}

//

func (c *Connection)GetConnId() uint32 {
	return c.connid
}

func (c *Connection)Start()  {
	fmt.Println("[Connection Start called..]")
	//TODo
	for  {
		//读取数据
		buf:=make([]byte,512)
		cnt,err:=c.tcpconn.Read(buf)
		if err!=nil {
			fmt.Printf("tcpconn.Read err",err)
			return
		}
		fmt.Println("Server <=== Content ",cnt,"data:",string(buf[:cnt]))
		//变成大写
		writBackInfo:=strings.ToUpper(string(buf[:cnt]))
		//服务器发送给客段
		cnt,err=c.tcpconn.Write([]byte(writBackInfo))
		fmt.Println("Server ===> Content",string(buf[:cnt]))
		//服务器发送给客户段
	}
}

func (c *Connection)Stop()  {
	fmt.Println("[Connection Start called..]")
}