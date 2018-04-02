package rediskey

import "gameTools/game/tools"

const (
	REDIS_KEY_SRV_LOGICLIST = "srv-LogicList"
	REDIS_KEY_SRV_SCENELIST = "srv-SceneList"
	REDIS_KEY_SRV_GATELIST = "srv-GateList"
	REDIS_KEY_PLAY_INFO = "playInfo-"
)

func GetPlayInfoKey(id int64) string{
	return REDIS_KEY_PLAY_INFO + tools.IntToString(id)
}