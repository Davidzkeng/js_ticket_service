package response

import (
	"github.com/gin-gonic/gin"
)

func JsonDefaultResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, 200, data)
}

func JsonResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
