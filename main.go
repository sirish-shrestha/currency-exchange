package main

import (
	"zumata-currency-exchange/app"
	"zumata-currency-exchange/config"
)

func main() {
	config := config.GetConfig()
	
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}