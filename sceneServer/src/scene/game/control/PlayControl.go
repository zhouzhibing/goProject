package control

import (
	"gameTools/game/tools/json"
	"scene/game/db/redis"
	"scene/game/db/mysql"
	"scene/game/msg"
	"gameTools/game/tools/protocol"
	"fmt"
	"gameTools/game/tools"
)

type PlayControl struct{
	userControl * UserControl
	sceneControl * SceneControl 		//当前所在的场景控制器
}

func NewPlayControl(userControl * UserControl) *PlayControl{

	this := new(PlayControl)
	this.userControl = userControl
	return this
}

func (this * PlayControl) GetId() int64 {
	return this.userControl.GetId()
}

func (this * PlayControl) GetUid() string {
	return this.userControl.GetUid()
}


func (this * PlayControl) Send(o *json.JSONObject) {

	o.Put(protocol.MSG_HANDLE_ID , this.GetId())
	o.Put("time" , tools.Millisecond())

	fmt.Println("Send ==> " , o)

	this.userControl.Send(o)
}

func (this *PlayControl) newPlayEnterScene(newPlayControl *PlayControl) {

	o := msg.GetNewPlayEnterScenePkg(newPlayControl.GetId() , newPlayControl.GetX() , newPlayControl.GetY() , newPlayControl.GetRobot())
	this.Send(o)
}

func (this *PlayControl) removePlayEnterScene(newPlayControl *PlayControl) {

	o := msg.GetRemovePlayExitScenePkg(newPlayControl.GetId())
	this.Send(o)
}

func (this * PlayControl) EnterScene(o * json.JSONObject) {

	sceneId := o.GetInt("toSceneId")

	this.sceneControl = SingleWorldManager().GetSceneControl(sceneId)

	this.sceneControl.EnterScene(this , o)

	this.userControl.SetSceneId(sceneId)

	this.Send(o)
}

func (this * PlayControl) SceneMove(o * json.JSONObject) {
	x := o.GetInt("x")
	y := o.GetInt("y")

	this.userControl.SetX(x)
	this.userControl.SetY(y)

	if this.sceneControl != nil{
		this.sceneControl.SceneMove(this , o)
	}

	//this.Send(o)
}


func (this *PlayControl) ExitScene(o *json.JSONObject) {

	this.sceneControl.ExitScene(this , o)

	user := this.userControl.user

	typeString := o.GetString("type")

	redis.UpdateGameObject(typeString , user)

	mysql.UpdateGameObject(user)

	toSceneId := o.GetInt("toSceneId")

	if !SingleWorldManager().ExistSceneId(toSceneId) {

		SingleWorldManager().RemoveUserControl(user.Id)
	}

}

func (this * PlayControl) GetX() int {	return this.userControl.GetX()}
func (this * PlayControl) GetY() int {	return this.userControl.GetY()}
func (this *PlayControl) GetRobot() bool {	return this.userControl.GetRobot() }



