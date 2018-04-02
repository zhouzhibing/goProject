package msg

import (
	"gameTools/game/tools/json"
	"gameTools/game/tools/protocol"
)

func GetEnterScene(o * json.JSONObject){

}

func GetPushSceneMovePkg( id int , x int , y int){

	other := json.NewJSONObject()
	other.Put("id" , id)
	other.Put("x" , x)
	other.Put("y" , y)
}


func GetNewPlayEnterScenePkg(id int64 , x int , y int , robot bool) * json.JSONObject{

	o := json.NewJSONObject()

	o.Put(protocol.MSG_PROTOCOL_NO , protocol.MSG_S2C_NEW_PLAY_ENTER_SCENE)
	o.Put("Id" , id)
	o.Put("X" , x)
	o.Put("Y" , y)
	o.Put("Robot" , robot)

	return o
}


func GetRemovePlayExitScenePkg(id int64) * json.JSONObject{

	o := json.NewJSONObject()

	o.Put(protocol.MSG_PROTOCOL_NO , protocol.MSG_S2C_NEW_PLAY_EXIT_SCENE)
	o.Put("Id" , id)
	return o
}



func GetSceneMovePkg(id int64 , x int , y int) * json.JSONObject{

	o := json.NewJSONObject()

	o.Put(protocol.MSG_PROTOCOL_NO , protocol.MSG_S2C_PLAY_SCENE_MOVE)
	o.Put("Id" , id)
	o.Put("X" , x)
	o.Put("Y" , y)

	return o
}








