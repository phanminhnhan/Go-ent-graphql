package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"go-ent-graphql/database"
	"go-ent-graphql/router"

)

func main() {
	e := echo.New()
	database.Init()	
	e.Use(middleware.Logger())
	router := router.Router_api{
		Echo: e,
	}
	router.Init_Router()
	e.Logger.Fatal(e.Start(":8080"))
}