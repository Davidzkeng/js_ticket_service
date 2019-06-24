package main

import (
	"fmt"
	"js_ticket_service/app"
)

func main() {
	fmt.Println("===========start js_ticket service")
	engine := app.Init
	app.Run(engine())

}
