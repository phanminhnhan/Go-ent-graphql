package middleware


import (
	"go-ent-graphql/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func UseJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc{
// 	config := middleware.JWTConfig{
// 		 Claims: &security.JwtCustomClaims{},
// 		 SigningKey: []byte(security.SecretKey),
// 		 TokenLookup: "query:token",

// 	}
// 	return middleware.JWTWithConfig(config)
// }

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &security.JwtCustomClaims{},
		SigningKey: []byte(security.SecretKey),
	}

	return middleware.JWTWithConfig(config)
}	