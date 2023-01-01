package controllers

import (
	"echo-auth/pkg/auth"
	"echo-auth/pkg/config"
	"echo-auth/pkg/models"
	"echo-auth/pkg/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func SignupController(c echo.Context) error {
	db := config.GetDatabase()
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal("Error bcrypt.GenerateFromPassword: ", err)
	}
	user.Password = string(hash)
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func LoginController(c echo.Context) error {
	user := models.User{}
	// Get the email / password from request
	var login struct {
		Username string
		Password string
	}
	if err := c.Bind(&login); err != nil {
		return err
	}
	db := config.GetDatabase()
	db.First(&user, "Username = ?", login.Username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		rs := utils.ResponseStruct{StatusCode: http.StatusForbidden, MessageEn: "Username/password invalid"}
		return rs.WriteToResponse(c, nil, "en")
	}

	// If password is correct, generate tokens and set cookies.
	token := auth.GenerateTokensAndSetCookies(&user, c)

	if err != nil {
		rs := utils.ResponseStruct{StatusCode: http.StatusUnauthorized, MessageEn: "Token is incorrect"}
		return rs.WriteToResponse(c, nil, "en")
	}
	rs := utils.ResponseStruct{StatusCode: http.StatusOK, MessageEn: "Login Successful"}
	return rs.WriteToResponse(c, token, "en")
}

func UserController(c echo.Context) error {
	// Gets user cookie.
	userCookie, _ := c.Cookie("user")
	db := config.GetDatabase()
	user := models.User{}
	db.Select("ID", "username", "firstname", "lastname").First(&user, "Username = ?", userCookie.Value)
	rs := utils.ResponseStruct{StatusCode: http.StatusOK, MessageEn: "User found"}
	return rs.WriteToResponse(c, user, "en")
}
