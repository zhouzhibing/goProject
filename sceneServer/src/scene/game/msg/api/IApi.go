package api

import (
	"gameTools/game/tools/net/sockets/socket"
	"gameTools/game/tools/json"
)

type IApi interface {
	Hanndle(connection * socket.Connection ,  protocolNo int , o * json.JSONObject)
}

