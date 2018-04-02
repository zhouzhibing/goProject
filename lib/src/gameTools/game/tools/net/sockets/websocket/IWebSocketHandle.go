package websocket

type IWebSocketHandle interface {

	Connection(connection * Connection)		//新的连接

	Handle(connection * Connection , msgObject interface{})			//处理连接逻辑

	Close(connection * Connection)				//连接关闭
}
