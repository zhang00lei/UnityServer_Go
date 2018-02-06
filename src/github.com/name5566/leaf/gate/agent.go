package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	//添加客户端断开连接事件
	AddSocketCloseEvent(f func(agent Agent))
	RemoveSocketCloseEvent(f func(agent Agent))
	UserData() interface{}
	SetUserData(data interface{})
}
