package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/middleware"
	"server/router"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	Router.Use(middleware.Cors())
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))
	Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))
	routerGroup := router.RouterGroupApp

	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth())
	adminGroup := Router.Group(global.Config.System.RouterPrefix)
	adminGroup.Use(middleware.JWTAuth()).Use(middleware.AdminAuth())
	{
		routerGroup.InitBaseRouter(publicGroup)
	}
	{
		routerGroup.InitUserRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitArticleRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitCommentRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitFeedbackRouter(privateGroup, publicGroup, adminGroup)
	}
	{
		routerGroup.InitImageRouter(adminGroup)
		routerGroup.InitAdvertisementRouter(adminGroup, publicGroup)
		routerGroup.InitFriendLinkRouter(adminGroup, publicGroup)
		routerGroup.InitWebsiteRouter(adminGroup, publicGroup)
		routerGroup.InitConfigRouter(adminGroup)
	}
	return Router
}
