package control

import (
	"sync"
	"fmt"
	"gameTools/game/tools/collect"
	"gameTools/game/tools"
)

type worldManager struct {

	qps * tools.TimeCount				//每秒请求数量
	tps * tools.TimeCount				//每秒处理数量

	sceneControls * SceneControls		//场景集合控制器
	userControlMap * collect.Map		//用户集合
}

var worldManagerObject * worldManager
var worldlock * sync.Mutex = tools.NewMutex()

func SingleWorldManager() *worldManager{

	if worldManagerObject != nil {
		return worldManagerObject
	}

	worldlock.Lock()
	defer worldlock.Unlock()

	if worldManagerObject != nil {
		return worldManagerObject
	}

	if worldManagerObject == nil {
		worldManagerObject = new(worldManager)
		worldManagerObject.init()

		return worldManagerObject
	}

	return worldManagerObject
}

func (this * worldManager) init(){
	this.userControlMap = collect.NewMap()
	this.sceneControls = NewSceneControls()
	this.qps = tools.NewTimeCount(1000)
	this.tps = tools.NewTimeCount(1000)
}


func (this * worldManager) Start(){
	fmt.Println("WorldManager Start")
}

func (this *worldManager) AddUserControl(userControl * UserControl) {
	this.userControlMap.Put(userControl.GetId() , userControl)
	//fmt.Println("WorldManager AddUserControl Size : " , this.userControlMap.Value())
}

func (this *worldManager) GetUserControl(uId int64) * UserControl{
	userControl := this.userControlMap.Get(uId)
	if userControl == nil{
		return nil
	}
	return userControl.(* UserControl)
}

func (this *worldManager) RemoveUserControl(uId int64) {
	this.userControlMap.Remove(uId)
	//fmt.Println("WorldManager RemoveUserControl Size : " , this.userControlMap.Value() )
}

func (this *worldManager) GetSceneControl(sceneId int) * SceneControl {
	return this.sceneControls.GetSceneControl(sceneId)
}


func (this *worldManager) ExistSceneId(sceneId int) bool{
	return this.sceneControls.ExistSceneId(sceneId)
}

func (this *worldManager) UserCount() int {	return this.userControlMap.Size()}
func (this * worldManager) Qps(){	this.qps.Count()}
func (this * worldManager) Tps(){	this.tps.Count()}
func (this * worldManager) GetQpsCount() int64 {	return this.qps.GetCount()		}
func (this * worldManager) GetTpsCount() int64 {		return this.tps.GetCount()		}