package control

import (
	"gameTools/game/tools/collect"
	"gameTools/game/tools/config"
	"gameTools/game/tools"
	"strings"
	"fmt"
)

type SceneControls struct {
	sceneControlMap * collect.Map
}

func NewSceneControls() * SceneControls{
	this := new(SceneControls)
	this.init()
	return this
}

func (this * SceneControls) init(){
	sceneIdString := config.GetSingleConfigValue("server.sceneId")

	sceneIdArray := strings.Split(sceneIdString , ",")

	this.sceneControlMap = collect.NewMap()

	for i := 0 ; i < len(sceneIdArray) ; i++{

		sceneId := tools.StringToInt(sceneIdArray[i])

		this.sceneControlMap.Put(sceneId , NewSceneControl(sceneId))
	}


	fmt.Println("SceneId By : " , sceneIdArray)
}

func (this * SceneControls) EnterScene(){

}

func (this *SceneControls) GetSceneControl(sceneId int) * SceneControl{
	sceneObject := this.sceneControlMap.Get(sceneId)
	if sceneObject == nil{
		return nil
	}
	return sceneObject.( * SceneControl)
}

func (this *SceneControls) ExistSceneId(sceneId int) bool {
	return this.sceneControlMap.ExistKey(sceneId)
}
