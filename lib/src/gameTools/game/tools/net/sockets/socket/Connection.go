package socket

import (
	"net"
	"uuid-master"
)

//socket 的连接对象
type Connection struct {

	conn net.Conn							//net 连接对象
	id string
	deencode IDeEncode						//对应的编码解码器
	attachment interface{}					//附件对象，用于连接绑定对象
	close bool

}

func NewConnection(conn net.Conn , deencode IDeEncode ) * Connection{
	this := new(Connection)

	this.conn = conn
	this.id = uuid.Rand().Hex()
	this.deencode = deencode

	return this
}

func (this * Connection) Read(b []byte) (n int, err error){
	return this.conn.Read(b)
}

func (this * Connection) Write(message interface{}){
	this.Send(message)
}

func (this * Connection) Send(message interface{}) error{

	buffer := this.deencode.Encode(message)

	_ , err := this.conn.Write(buffer)
	//fmt.Println("send n : " , n )

	return err
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
	this.close = true
	this.conn.Close()
}

func (this *Connection) IsOpen() bool{
	if this.close {
		return false
	}else{
		return true
	}
}

func (this * Connection) LocalAddr() net.Addr{
	return this.conn.LocalAddr()
}

func (this * Connection) RemoteAddr() net.Addr{
	return this.conn.RemoteAddr()
}



