package handler

import (
	"net/http"
	"io"
	// "mime/multipart"
	"os"
	"github.com/labstack/echo/v4"
	"go-ent-graphql/model"
	"go-ent-graphql/database"
	"go-ent-graphql/ent"
	"go-ent-graphql/security"
	"go-ent-graphql/ent/user"
	// "context"
	"strconv"
	"errors"
	"log"
	//"fmt"
	"github.com/golang-jwt/jwt"
)


func CreateUser(c echo.Context)error{
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	age := c.FormValue("age")

	file, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}	
	paseAge, _ := strconv.Atoi(age)
	hash := security.HashAndSalt([]byte(password))
	createdUser, err := database.EntClient.User.
      Create().                      
      SetName(name).
      SetEmail(email).
      SetPassword(hash).
	  SetAge(paseAge).
	  SetAvatar(file.Filename).
      Save(c.Request().Context())  
	token, err := security.Gentoken(createdUser.ID.String(), name, email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}
	data := model.Authen_Response{
		Token : token["access_token"],
		Refresh_token: token["refresh_token"],
		User: createdUser,
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "done",
		Data: data,
	}) 
}

func GetUser(c echo.Context) error{
	user_req := new(struct{
		Email 		string      	`json:"email"`
		Password 	string	  		`json:"password"`
	})
	err := c.Bind(&user_req)
	if err != nil{
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}
	user, err := database.EntClient.User.
		Query().
		Where(user.Email(user_req.Email)).
		Only(c.Request().Context())
	if err != nil{
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}
	isTheSame := security.ComparePasswords(user.Password, []byte(user_req.Password))
	if !isTheSame {
		// log.Error()
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "dang nhap that bai",
			Data:       nil,
		})
	}
	token, err := security.Gentoken(user.ID.String(), user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}
	data := model.Authen_Response{
		Token : token["access_token"],
		Refresh_token: token["refresh_token"],
		User: user,
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "done",
		Data: data,
	}) 
}

func GetUsers(e echo.Context) error{
	pageStr := e.QueryParam("page")
	page, _ := strconv.Atoi(pageStr)
	limitStr := e.QueryParam("limit")
	limit, _ := strconv.Atoi(limitStr)

	if page < 0 || limit < 0 {
		return e.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "khong hop le",
			Data:       nil,
		})
	}
	offset := (page-1)*limit
	users, err := database.EntClient.User.
		Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Asc("created_at")).
		All(e.Request().Context())
	if !errors.Is(err, nil) {
		log.Fatalf("Error: failed quering users %v\n", err)
	}
	return e.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "done",
		Data: users,
	}) 
}


func UserProfile(c echo.Context)error{
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*security.JwtCustomClaims)
	userEmail := claims.Email
	user, err := database.EntClient.User.
		Query().
		Where(user.Email(userEmail)).
		Only(c.Request().Context())
	if err != nil{
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.JSON(http.StatusInternalServerError, model.Response{
		StatusCode: 200,
		Message: "done",
		Data: user,
	})
}


func TestUploadFile(c echo.Context) error{
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.JSON(http.StatusInternalServerError, model.Response{
		StatusCode: 200,
		Message: "done",
		Data: file,
	})
}