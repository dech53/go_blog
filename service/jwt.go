package service

import (
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"server/global"
	"server/model/database"
	"server/utils"
)

type JwtService struct{}

func (service *JwtService) SetRedisJWT(jwt string, uuid uuid.UUID) error {
	dr, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		return err
	}
	return global.Redis.Set(uuid.String(), jwt, dr).Err()
}

func (service *JwtService) GetRedisJWT(uuid uuid.UUID) (string, error) {
	return global.Redis.Get(uuid.String()).Result()
}

func (service *JwtService) JoinInBlackList(jwtList database.JwtBlacklist) error {
	if err := global.DB.Create(&jwtList).Error; err != nil {
		return err
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return nil
}

func (service *JwtService) IsInBlackList(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func LoadAll() {
	var data []string
	if err := global.DB.Model(&database.JwtBlacklist{}).Pluck("jwt", &data).Error; err != nil {
		global.Log.Error("Failed to load JWT blacklist from the database", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
