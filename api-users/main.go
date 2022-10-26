package main

import (
	"api-users/config"
	"api-users/database"
	"api-users/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.GetConfig()

	e := echo.New()

	//run database
	database.ConnectDB()

	//routes
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(conf.Port))
}
