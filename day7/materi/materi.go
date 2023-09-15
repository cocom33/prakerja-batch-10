package main

import (
	"prakerja/day7/materi/configs"
	"prakerja/day7/materi/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	
	e.Start(":8000")
}