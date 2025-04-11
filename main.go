package main

import (
	"server/core"
	"server/flag"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()
	global.MinioClient = initialize.InitMinio()

	defer global.Redis.Close()
	flag.InitFlag()
	initialize.InitCron()

	core.RunServer()
}
