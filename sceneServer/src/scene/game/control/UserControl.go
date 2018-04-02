package control

import (
	"gameTools/game/tools/collect"
	"scene/game/db/entity"
	"gameTools/game/tools/net/sockets/socket"
	"gameTools/game/tools/json"
)

type UserControl struct {

	user * entity.User
	playContorlMap * collect.Map
	connection * socket.Connection
}

func NewUserControl(user * entity.User , connection * socket.Connection) *UserControl{

	this := new(UserControl)
	this.user = user
	this.connection = connection

	this.init()

	return this
}

func (this * UserControl) init(){

	if this.playContorlMap == nil{
		this.playContorlMap = collect.NewMap()
	}

	playControl := NewPlayControl(this)

	this.playContorlMap.Put(playControl.GetId() , playControl)
}

func (this * UserControl) GetId() int64{
	return this.user.Id
}
func (this * UserControl) GetUid() string {
	return this.user.UId
}


func (this * UserControl) GetPlayControl() * PlayControl{

	for _ , playContorl := range this.playContorlMap.Value(){
		return playContorl.(* PlayControl)
	}
	return nil
}

func (this * UserControl) Send(o *json.JSONObject){
	err := this.connection.Send(o.ToJSONString())
	if err != nil{
		panic(err)
	}
	SingleWorldManager().Tps()
}


func (this * UserControl) SetSceneId(sceneId int) {	this.user.Scene_id = sceneId}
func (this * UserControl) GetX() int {	return this.user.X}
func (this * UserControl) GetY() int {	return this.user.Y}
func (this * UserControl) SetX(x int){	this.user.X = x}
func (this * UserControl) SetY(y int) {	this.user.Y = y}
func (this * UserControl) GetSceneId() int {	return this.user.Scene_id}
func (this * UserControl) GetRobot() bool {	return this.user.Robot}
