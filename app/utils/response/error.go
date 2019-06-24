package response

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func NotFound(c *gin.Context) {
	JsonResponse(c, 404, ErrorMessage{-1,"RESOURCE_NOT_FOUND", "Not Found"})
}

func InvalidRequest(c *gin.Context, name string, message string) {

	if name == "" {
		name = "INVALID_REQUEST"
	}

	if message == "" {
		message = "Invalid Request"
	}

	JsonResponse(c, 400, ErrorMessage{-1,name, message})
}

func InternalError(c *gin.Context) {
	JsonResponse(c, 500, ErrorMessage{-1,"INTERNAL_SERVER_ERROR", "Internal Server Error"})
}
