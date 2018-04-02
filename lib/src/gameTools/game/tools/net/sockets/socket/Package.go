package socket

type Package struct {
	conn * Connection
	msgObject interface{}
}

func NewPackage(conn *Connection , msgObject interface{}) * Package{
	this := new(Package)
	this.conn = conn
	this.msgObject = msgObject
	return this
}

func (this * Package) GetConnection() *Connection{
	return this.conn
}

func (this * Package) GetMsgObject() interface{}{
	return this.msgObject
}