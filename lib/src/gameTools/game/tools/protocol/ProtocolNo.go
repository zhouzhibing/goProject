package protocol

const (
	MSG_PROTOCOL_NO = "protocolNo"
	MSG_HANDLE_ID = "handleId"

	//
	MSG_C2S_ENTER_SCENE = 100001
	MSG_C2S_SCENE_MOVE = 100002					//场景内有人移动
	MSG_S2S_EXIT_SCENE = 100003					//退出离开游戏


	MSG_S2C_NEW_PLAY_ENTER_SCENE = 900001		//新角色进入场景推送
	MSG_S2C_NEW_PLAY_EXIT_SCENE = 900002		//有角色离开场景推送
	MSG_S2C_PLAY_SCENE_MOVE = 900003			//场景内有人移动推送

)


func GetMode(protocolNo int) int {
	return int (protocolNo / 10000);
}
