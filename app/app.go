package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"js_ticket_service/app/routes"
	"js_ticket_service/app/utils/log"
	"js_ticket_service/config"
	"os"
)

func Init() *gin.Engine {
	err := config.LoadConfigFromToml()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	engine := gin.New()

	setDefaultConfig()
	routes.New(engine)
	logger.NewLogger(config.Cfg.GetString("log.path"), config.Cfg.GetString("log.level"), config.Cfg.GetBool("log.isDebug"))


	return engine

}

func setDefaultConfig() {
	if config.Cfg.GetString("application.env") == "dev" {
		gin.SetMode(gin.DebugMode)
		return
	}
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
	gin.DefaultWriter = ioutil.Discard
}

func Run(e *gin.Engine) {
	e.Run(config.Cfg.GetString("const.addr"))
}
