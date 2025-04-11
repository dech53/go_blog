package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/config"
	"server/global"
	"server/model/response"
)

type ConfigApi struct {
}

func (configApi *ConfigApi) GetWebsite(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

func (configApi *ConfigApi) UpdateWebsite(c *gin.Context) {
	var req config.Website
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateWebsite(req)
	if err != nil {
		global.Log.Error("Failed to update website:", zap.Error(err))
		response.FailWithMessage("Failed to update website", c)
		return
	}
	response.OkWithMessage("Successfully updated website", c)
}

func (configApi *ConfigApi) GetSystem(c *gin.Context) {
	response.OkWithData(global.Config.System, c)
}

func (configApi *ConfigApi) UpdateSystem(c *gin.Context) {
	var req config.System
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateSystem(req)
	if err != nil {
		global.Log.Error("Failed to update system:", zap.Error(err))
		response.FailWithMessage("Failed to update system", c)
		return
	}
	response.OkWithMessage("Successfully updated system", c)
}

func (configApi *ConfigApi) GetEmail(c *gin.Context) {
	response.OkWithData(global.Config.Email, c)
}

func (configApi *ConfigApi) UpdateEmail(c *gin.Context) {
	var req config.Email
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateEmail(req)
	if err != nil {
		global.Log.Error("Failed to update email:", zap.Error(err))
		response.FailWithMessage("Failed to update email", c)
		return
	}
	response.OkWithMessage("Successfully updated email", c)
}

func (configApi *ConfigApi) GetJwt(c *gin.Context) {
	response.OkWithData(global.Config.Jwt, c)
}

func (configApi *ConfigApi) UpdateJwt(c *gin.Context) {
	var req config.Jwt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateJwt(req)
	if err != nil {
		global.Log.Error("Failed to update jwt:", zap.Error(err))
		response.FailWithMessage("Failed to update jwt", c)
		return
	}
	response.OkWithMessage("Successfully updated jwt", c)
}

func (configApi *ConfigApi) GetGaode(c *gin.Context) {
	response.OkWithData(global.Config.Gaode, c)
}

func (configApi *ConfigApi) UpdateGaode(c *gin.Context) {
	var req config.Gaode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateGaode(req)
	if err != nil {
		global.Log.Error("Failed to update gaode:", zap.Error(err))
		response.FailWithMessage("Failed to update gaode", c)
		return
	}
	response.OkWithMessage("Successfully updated gaode", c)
}
