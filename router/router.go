package router

import (
	"github.com/labstack/echo/v4"
	"go-ent-graphql/handler"
	"go-ent-graphql/middleware"
)

type Router_api struct {
	Echo *echo.Echo
}

func (r *Router_api)Init_Router(){
	r.Echo.GET("/check", handler.Health_check)
	r.Echo.POST("/upload", handler.TestUploadFile)
	r.Echo.POST("/sign-up", handler.CreateUser)
	r.Echo.POST("/sign-in", handler.GetUser)
	r.Echo.GET("/users", handler.GetUsers)
	r.Echo.GET("/me", handler.UserProfile, middleware.JWTMiddleware())
}