package core

import (
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"server/service"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()

	service.LoadAll()

	s := initServer(addr, Router)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
