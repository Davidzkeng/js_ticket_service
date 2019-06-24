package routes

import (
	"github.com/gin-gonic/gin"
	"js_ticket_service/app/http/controllers"
)

func New(e *gin.Engine) {
	e.Use(gin.Logger())
	e.POST("/get_jsticket", controllers.GetTicket)
}
