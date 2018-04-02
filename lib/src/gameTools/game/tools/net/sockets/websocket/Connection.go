package websocket

import (
	"net"
	"golang.org/x/net/websocket"
	"uuid-master"
)

//socket 的连接对象
type Connection struct {

	conn * websocket.Conn							//net 连接对象
	id string
	attachment interface{}					//附件对象，用于连接绑定对象

}

func NewConnection(conn*  websocket.Conn) * Connection{
	this := new(Connection)
	this.conn = conn
	this.id = uuid.Rand().Hex()

	return this
}

func (this * Connection) RemoteAddr () net.Addr{
	return this.conn.RemoteAddr()
}

func (this * Connection) Read(b []byte) (n int, err error){
	return this.conn.Read(b)
}

func (this * Connection) Write(message interface{}){
	this.Send(message)
}

func (this * Connection) Send(message interface{}) error{
	return websocket.Message.Send(this.conn, message)
}

func (this * Connection) SetAttachment(attackment interface{}){
	this.attachment = attackment
}

func (this * Connection) GetAttachment() interface{} {
	return this.attachment
}

func (this * Connection) GetId() string  {
	return this.id
}

func (this *Connection) Close() {
	this.conn.Close()
}