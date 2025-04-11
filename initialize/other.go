package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"os"
	"server/global"
	"server/utils"
)

func OtherInit() {
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)
}
