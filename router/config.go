package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type ConfigRouter struct {
}

func (c *ConfigRouter) InitConfigRouter(Router *gin.RouterGroup) {
	configRouter := Router.Group("config")

	configApi := api.ApiGroupApp.ConfigApi
	{
		configRouter.GET("website", configApi.GetWebsite)
		configRouter.PUT("website", configApi.UpdateWebsite)
		configRouter.GET("system", configApi.GetSystem)
		configRouter.PUT("system", configApi.UpdateSystem)
		configRouter.GET("email", configApi.GetEmail)
		configRouter.PUT("email", configApi.UpdateEmail)
		configRouter.GET("jwt", configApi.GetJwt)
		configRouter.PUT("jwt", configApi.UpdateJwt)
		configRouter.GET("gaode", configApi.GetGaode)
		configRouter.PUT("gaode", configApi.UpdateGaode)
	}
}
