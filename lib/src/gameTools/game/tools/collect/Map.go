package collect

import (
	"fmt"
	"sync"
	"gameTools/game/tools"
)

type Map struct {
	mapValue map[interface{}]interface{}
	lock * sync.Mutex
	rwLock * sync.RWMutex
}

func NewMap() * Map{
	this := new(Map)
	this.init()
	return this
}

func (this *Map ) init(){
	if this.lock == nil{
		this.lock = tools.NewMutex()
		//this.rwLock = tools.NewRwMutex()
	}
	//fmt.Println("Map init !")
}

func (this *Map ) Put(key interface{} , value interface{}){
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.mapValue == nil{
		this.mapValue = make(map[interface{}]interface{} , 10 );
	}

	this.mapValue[key] = value
}

func (this * Map) Get(key interface{})interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()


	if this.mapValue == nil{
		return nil
	}
	return this.mapValue[key]
}

func (this * Map) Remove(key interface{}){
	this.lock.Lock()
	defer this.lock.Unlock()

	delete(this.mapValue,key)
}

func (this * Map) Value() map[interface{}]interface{}{
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.mapValue;
}

func (this * Map) Pool(exec func(key interface{} , value interface{})){
	this.lock.Lock()
	defer this.lock.Unlock()

	for key , value := range this.mapValue{
		exec(key , value)
	}
}

func (this * Map) Print(){
	fmt.Println(this.mapValue)
}

func (this *Map) Size() int {
	size := len(this.mapValue)
	return size
}

func (this *Map) ExistKey(key interface{}) bool {
	for inKey , _ := range this.mapValue{
		if inKey == key{
			return true
		}
	}
	return false
}