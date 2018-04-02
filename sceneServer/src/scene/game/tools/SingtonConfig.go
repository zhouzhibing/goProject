package tools

import (
	"sync"
	"gameTools/game/tools/config"
)

type ServerConfig struct {
	config * config.Config
}

var serverConfig * ServerConfig
var lock sync.Mutex

func SingtonConfig(path string) * ServerConfig{
	if serverConfig == nil{
		lock.Lock()
		if serverConfig == nil{
			serverConfig = new(ServerConfig)
		}
		lock.Unlock()
	}
	serverConfig.Init(path)
	return serverConfig
}

func (this * ServerConfig) Init(path string){
	this.config = config.NewConfig(path)
}

func (this *ServerConfig) GetValue(key string)  string {
	return this.config.GetValue(key)
}




