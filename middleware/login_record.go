package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/zap"
	"server/global"
	"server/model/database"
	"server/service"
)

func LoginRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		go func() {
			gaodeService := service.ServiceGroupApp.GaodeService
			var userID uint
			var address string
			ip := c.ClientIP()
			userAgent := c.Request.UserAgent()

			if value, exists := c.Get("user_id"); exists {
				if id, ok := value.(uint); ok {
					userID = id
				}
			}

			address = getAddressFromIP(ip, gaodeService)

			os, deviceInfo, browserInfo := parseUserAgent(userAgent)

			login := database.Login{
				UserID:      userID,
				IP:          ip,
				Address:     address,
				OS:          os,
				DeviceInfo:  deviceInfo,
				BrowserInfo: browserInfo,
				Status:      c.Writer.Status(),
			}

			if err := global.DB.Create(&login).Error; err != nil {
				global.Log.Error("Failed to record login", zap.Error(err))
			}
		}()
	}
}

func getAddressFromIP(ip string, gaodeService service.GaodeService) string {
	res, err := gaodeService.GetLocationByIP(ip)
	if err != nil || res.Province == "" {
		return "未知"
	}
	if res.City != "" && res.Province != res.City {
		return res.Province + "-" + res.City
	}
	return res.Province
}

func parseUserAgent(userAgent string) (os, deviceInfo, browserInfo string) {
	os = userAgent
	deviceInfo = userAgent
	browserInfo = userAgent

	parser := uaparser.NewFromSaved()
	cli := parser.Parse(userAgent)
	os = cli.Os.Family
	deviceInfo = cli.Device.Family
	browserInfo = cli.UserAgent.Family

	return
}
