package redis

import (
	"gameTools/game/tools/db/redis"
	config "gameTools/game/tools/config"
	"errors"
	"gameTools/game/tools/db/rediskey"
	"login/db/entity"
	"gameTools/game/tools/json"
)

var redisPool *redis.RedisPool

func RedisControlInit() {

	if redisPool != nil{
		panic(errors.New("redis pool already register !"))
	}

	address := config.GetSingleConfigValue("redis.address")
	pass := config.GetSingleConfigValue("redis.pass")

	redisPool = redis.NewRedisPool(address , pass)
}

func GetMinGateServer() interface {}{
	return redisPool.Do("hgetAll", rediskey.REDIS_KEY_SRV_GATELIST)
}

func UpdateGameObject(user* entity.User){

	playKey := rediskey.GetPlayInfoKey(user.Id)
	jsonString := json.JSONParseString(user)

	redisPool.Do("hset"  , playKey , "User" , jsonString )
}


