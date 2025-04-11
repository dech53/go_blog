package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"time"
)

type WebsiteApi struct {
}

func (website *WebsiteApi) WebsiteLogo(c *gin.Context) {
	if global.Config.Website.Logo != "" {
		c.Redirect(http.StatusMovedPermanently, global.Config.Website.Logo)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/image/logo.png")
	}
}

func (website *WebsiteApi) WebsiteTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"title": global.Config.Website.Title})
}

func (website *WebsiteApi) WebsiteInfo(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

func (website *WebsiteApi) WebsiteCarousel(c *gin.Context) {
	urls := websiteService.WebsiteCarousel()
	response.OkWithData(urls, c)
}

func (website *WebsiteApi) WebsiteNews(c *gin.Context) {
	sourceStr := c.Query("source")
	hotSearchData, err := websiteService.WebsiteNews(sourceStr)
	if err != nil {
		global.Log.Error("Failed to get news:", zap.Error(err))
		response.FailWithMessage("Failed to get news", c)
		return
	}
	response.OkWithData(hotSearchData, c)
}

func (website *WebsiteApi) WebsiteCalendar(c *gin.Context) {
	dateStr := time.Now().Format("2006/0102")
	calendar, err := websiteService.WebsiteCalendar(dateStr)
	if err != nil {
		global.Log.Error("Failed to get calendar:", zap.Error(err))
		response.FailWithMessage("Failed to get calendar", c)
		return
	}
	response.OkWithData(calendar, c)
}

func (website *WebsiteApi) WebsiteFooterLink(c *gin.Context) {
	footerLinks := websiteService.WebsiteFooterLink()
	response.OkWithData(footerLinks, c)
}

func (website *WebsiteApi) WebsiteAddCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteAddCarousel(req)

	if err != nil {
		global.Log.Error("Failed to add carousel:", zap.Error(err))
		response.FailWithMessage("Failed to add carousel", c)
		return
	}
	response.OkWithMessage("Successfully added carousel", c)
}

func (website *WebsiteApi) WebsiteCancelCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteCancelCarousel(req)

	if err != nil {
		global.Log.Error("Failed to cancel carousel:", zap.Error(err))
		response.FailWithMessage("Failed to cancel carousel", c)
		return
	}
	response.OkWithMessage("Successfully canceled carousel", c)
}

func (website *WebsiteApi) WebsiteCreateFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteCreateFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to create footer link:", zap.Error(err))
		response.FailWithMessage("Failed to create footer link", c)
		return
	}
	response.OkWithMessage("Successfully created footer link", c)
}

func (website *WebsiteApi) WebsiteDeleteFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteDeleteFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to delete footer link:", zap.Error(err))
		response.FailWithMessage("Failed to delete footer link", c)
		return
	}
	response.OkWithMessage("Successfully deleted footer link", c)
}
