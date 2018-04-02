package mysql

import (
	"errors"
	"gameTools/game/tools/config"
	mysqltool "gameTools/game/tools/db/mysql"
	"login/db/entity"
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

func SelectUserByUid(uId string) * entity.User{
	var user *entity.User

	list := mysql.Query("select id , uId , scene_id , x , y from `User` where uId = ? " , &entity.User{} , uId)
	if list.Size() == 0 {
		user = createDefaultUser(uId)
		mysql.InsertObject(user)
	}else{
		user = list.Get(0).(* entity.User)
	}

	return user
}


func createDefaultUser(uId string) * entity.User{
	user := new(entity.User)
	user.Scene_id = 1111
	user.UId = uId
	user.X = tools.ToolsRandom(10 , 600 )
	user.Y = tools.ToolsRandom(10 , 600 )
	return user
}