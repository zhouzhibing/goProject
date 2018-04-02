package control

import (
	"gameTools/game/tools/collect"
	"gameTools/game/tools/json"
	"scene/game/msg"
)

type SceneControl struct {
	sceneId int
	playControlMap * collect.Map
}

func NewSceneControl(sceneId int) * SceneControl{
	this := new(SceneControl)
	this.sceneId = sceneId
	this.init()

	return this
}

func (this * SceneControl ) init(){
	this.playControlMap = collect.NewMap()
}

//有人进入场景
func (this * SceneControl) EnterScene(playControl * PlayControl , o * json.JSONObject){

	if this.playControlMap.Size() > 0 {

		otherPlayInfoList := collect.NewArrayList()

		this.playControlMap.Pool(func(index interface{}, value interface{}) {

			otherControl := value.( *PlayControl)

			if otherControl.GetId() != playControl.GetId() {

				other := json.NewJSONObject()
				other.Put("Id" , otherControl.GetId())
				other.Put("X" , otherControl.GetX())
				other.Put("Y" , otherControl.GetY())
				other.Put("Robot" , otherControl.GetRobot())

				otherPlayInfoList.Add(other.Value())

				//fmt.Println("otherControl.GetId() = " , otherControl.GetId() , " playControl.GetId() = " , playControl.GetId())

				otherControl.newPlayEnterScene(playControl)
			}
		})

		//fmt.Println("otherPlayInfoList: " , json.JSONParseString(otherPlayInfoList.Value()))

		o.Put("otherPlayList" , otherPlayInfoList.Value())

	}

	o.Put("selfPlay" , playControl.userControl.user)

	this.playControlMap.Put(playControl.GetId() , playControl)

	//fmt.Println("SceneControl EnterScene Size : " , this.playControlMap.Value())
}

//有人退出场景
func (this * SceneControl) ExitScene(playControl * PlayControl , o * json.JSONObject){

	this.playControlMap.Remove(playControl.GetId())

	this.playControlMap.Pool(func(index interface{}, value interface{}) {

		otherControl := value.( *PlayControl)

		otherControl.removePlayEnterScene(playControl)
	})

	//fmt.Println("SceneControl ExitScene Size : " , this.playControlMap.Value())
}

//场景移动
func (this * SceneControl) SceneMove(playControl *PlayControl , o * json.JSONObject){

	//fmt.Println("this.playControlMap = " , this.playControlMap)

	if this.playControlMap == nil{
		return
	}

	pushMovePkg := msg.GetSceneMovePkg(playControl.GetId() ,playControl.GetX() , playControl.GetY())

	this.playControlMap.Pool(func(index interface{}, value interface{}) {
		otherControl := value.( *PlayControl)

		if otherControl.GetId() != playControl.GetId() {			//排除自己
			otherControl.Send(pushMovePkg)
		}
	})
}


func (this * SceneControl) broadcast(o * json.JSONObject){

	this.playControlMap.Pool(func(index interface{}, value interface{}) {
		otherControl := value.( *PlayControl)
		otherControl.Send(o)
	})
}

func (this * SceneControl) GetSceneId() int{
	return this.sceneId
}