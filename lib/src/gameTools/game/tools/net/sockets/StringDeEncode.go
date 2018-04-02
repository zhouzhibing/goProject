package sockets

import (
	"bytes"
	"encoding/binary"
)
const PROTOCAL_HEAD_LENGTH = 4			//协议头长度

type stringDeEncode struct {}

func NewStringDeEncode() * stringDeEncode{
	this := new(stringDeEncode)
	return this
}

func (this * stringDeEncode) Decode(buffer [] byte , readerChannel chan interface{}) [] byte{

	totalLength := len(buffer)		//总长度

	//fmt.Println("totalLength = " , totalLength)

	if totalLength <  PROTOCAL_HEAD_LENGTH {			//是否满足包头长度
		return nil
	}

	var remainBuffer [] byte				//余下的数据buff

	for index := 0 ; index < totalLength;  {

		bodyLength := bytesToInt(buffer[index : index + PROTOCAL_HEAD_LENGTH])

		//fmt.Println("bodyLength = " , bodyLength)

		if bodyLength <= 0{
			return nil
		}

		startIndex := index + PROTOCAL_HEAD_LENGTH		//包体开始索引
		endIndex := startIndex + bodyLength							//包体结束索引

		if endIndex <= totalLength{			//符合长度，还可以解出包体

			subBuffer := buffer[startIndex : endIndex]

			dataString := string(subBuffer)

			//fmt.Println("dataString = " , dataString)

			readerChannel <- dataString

			index = endIndex

		}else{

			remainBuffer = buffer[index:]

			index = totalLength

		}
	}

	return remainBuffer
}


func (this * stringDeEncode) Encode(content interface{}) [] byte{

	contentString := content.(string)
	length := len(contentString)

	buffer := new(bytes.Buffer)
	buffer.Write(intToBytes(length))
	buffer.WriteString(contentString)

	return buffer.Bytes()
}

//整形转换成字节
func IntToBytes(n int) []byte {
	return intToBytes(n)
}

func intToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}