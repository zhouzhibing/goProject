package collect

import (
	"fmt"
)

type ArrayList struct{
	 elementData [] interface{}
}

func NewArrayList() * ArrayList{
	arrayList := new(ArrayList)

	arrayList.init()
	return arrayList
}

func (this * ArrayList) init(){
	//fmt.Println("ArrayList init !")
}

func (this * ArrayList) BuildData(elementData [] interface{}){
	if len(this.elementData) > 0 {
		fmt.Println("ArrayList elementData already data !")
		return
	}
	this.elementData = 	elementData
}

func (this * ArrayList) Add( add interface{}){
	this.elementData = append(this.elementData , add)
}

func (this * ArrayList) Remove(removeIndex int){
	this.elementData = append(this.elementData[:removeIndex] , this.elementData[removeIndex + 1:]...)
}

func (this * ArrayList) Get(index int) interface{} {
	return this.elementData[index]
}

func (this * ArrayList) Value() [] interface{}{
	return this.elementData
}

func (this * ArrayList) Size() int {
	if this == nil{
		return 0
	}
	if this.elementData == nil{
		return 0
	}
	return len(this.elementData)

}

func (this * ArrayList) Pool(exec func (index interface{} , value interface{})){
	for index , value := range this.elementData{
		exec(index , value)
	}
}

func (this * ArrayList)String()  string{
	fmt.Println("String")
	return fmt.Sprint(this.elementData...)
}

func (this * ArrayList) Print(){
	fmt.Println(this.elementData)
}

