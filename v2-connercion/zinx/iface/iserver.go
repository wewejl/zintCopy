package iface


type IServer interface {
	Start() //
	Stop()
	Serve()//服务函数
}