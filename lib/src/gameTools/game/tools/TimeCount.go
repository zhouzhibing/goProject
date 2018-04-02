package tools

import (
	"sync"
)

type TimeCount struct {
	lock * sync.Mutex
	timeCycle int64					//指定循环时间戳

	time int64							//毫秒时间戳
	count int64							//每秒请求处理的数量

	nowTime int64
	gapTime int64
}

func NewTimeCount(timeCycle int64 ) * TimeCount{
	this := new(TimeCount)
	this.timeCycle = timeCycle
	this.lock = NewMutex()
	return this
}

func (this * TimeCount) Count(){
	this.lock.Lock()
	defer this.lock.Unlock()

	this.nowTime = Millisecond()
	this.gapTime = this.nowTime - this.time

	if this.gapTime >= this.timeCycle{
		this.count = 0
	}

	this.count ++
	this.time = this.nowTime
}

func (this * TimeCount) GetCount() int64 {
	return this.count
}

