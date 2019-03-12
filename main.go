package main

import (
	"github.com/sirish-shrestha/zumata-currency-exchange/app"
	"github.com/sirish-shrestha/zumata-currency-exchange/config"
)

func main() {
	config := config.GetConfig()
	
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}