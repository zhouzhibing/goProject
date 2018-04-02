package sockets

//socket 的编码解码接口
type IDeEncode interface {

	Decode(buffer [] byte, readerChannel chan interface{}) [] byte		//解码
	Encode(content interface{})	[] byte														//编码

}



