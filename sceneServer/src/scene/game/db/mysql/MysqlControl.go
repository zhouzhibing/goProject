package mysql

import (
	"errors"
	"gameTools/game/tools/config"
	mysqltool "gameTools/game/tools/db/mysql"
	"scene/game/db/entity"
	"gameTools/game/tools"
)

var mysql * mysqltool.Mysql

func MysqlControlInit() {

	if mysql != nil{
		panic(errors.New("mysql already register !"))
	}

	address := config.GetSingleConfigValue("mysql.address")
	port := config.GetSingleConfigValue("mysql.port")
	dbname := config.GetSingleConfigValue("mysql.dbname")
	username := config.GetSingleConfigValue("mysql.username")
	pass := config.GetSingleConfigValue("mysql.pass")

	mysql = mysqltool.NewMysql(address , port , username ,pass , dbname )
}


func UpdateGameObject(user* entity.User){
	mysql.UpdateObject(user)
}


func SelectUserByUid(id int64) * entity.User{
	var user *entity.User

	list := mysql.Query("select id , uId , scene_id , x , y from `User` where uId = ? " , &entity.User{} , id)
	if list.Size() == 0 {
		user = createDefaultUser(id)
		mysql.InsertObject(user)
	}else{
		user = list.Get(0).(* entity.User)
	}

	return user
}


func createDefaultUser(id int64) * entity.User{

	user := new(entity.User)
	user.Scene_id = 1111
	user.Id = id
	user.X = tools.ToolsRandom(10 , 600 )
	user.Y = tools.ToolsRandom(10 , 600 )
	return user
}