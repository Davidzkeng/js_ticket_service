package controllers

import (
	"github.com/gin-gonic/gin"
	"js_ticket_service/app/utils/log"
	"js_ticket_service/app/utils/response"
	wx2 "js_ticket_service/app/utils/wx"
	"sync"
)

var wx wx2.Wx

func GetTicket(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		response.InternalError(c)
		return
	}

	accessToken := c.Request.PostForm["access_token"]
	appId := c.Request.PostForm["appid"]
	url := c.Request.PostForm["url"]

	if len(accessToken) == 0 {
		response.InvalidRequest(c, "", "missing access_token param")
		return
	}

	if len(appId) == 0 {
		response.InvalidRequest(c, "", "missing appid param")
		return
	}

	if len(url) == 0 {
		response.InvalidRequest(c, "", "missing url param")
		return
	}

	wx := wx2.Wx{
		AppId:       appId[0],
		AccessToken: accessToken[0],
		Lock:        sync.Mutex{},
	}
	res, err := wx.GetJsConfig(url[0])
	if err != nil {
		response.InvalidRequest(c, "", err.Error())
		logger.Logger.Error(err.Error())
		return
	}
	response.JsonResponse(c, 200, res)

}
