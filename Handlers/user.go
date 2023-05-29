package Handlers

import (
	"errors"
	"main/DB"
	"main/helpers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	var role, user string
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	} else {
		role, user, _ = helpers.FindRole(c)
	}
	if role == "user" {
		c.HTML(http.StatusOK, "dayout.html", user)
	} else {
		c.Redirect(http.StatusFound, "/admin")
	}
}

func SignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPost(c *gin.Context) {
	userName := c.Request.FormValue("UserName")
	userEmail := c.Request.FormValue("Emailid")
	password := c.Request.FormValue("Password")
	user := models.User{Name: userName, Email: userEmail, Password: password}
	result := DB.Db.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	c.Redirect(http.StatusFound, "/login")

}

func LoginHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func LoginPost(c *gin.Context) {
	var e models.Invalid
	newmail := c.Request.FormValue("Email")
	newpassword := c.Request.FormValue("Password")
	user := models.User{Email: newmail, Password: newpassword}
	result := DB.Db.Where(&models.User{Email: user.Email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		e.Errmail = "invalid email id"
		c.HTML(http.StatusOK, "login.html", e)
	} else if user.Password != newpassword {
		e.Errpass = "invalid password"
		c.HTML(http.StatusOK, "login.html", e)
	} else if user.Role == "user" {
		helpers.CreateToken(user, c)
		c.Redirect(http.StatusFound, "/")
	} else {
		helpers.CreateToken(user, c)
		c.Redirect(http.StatusFound, "/admin")
	}
}

func HomeHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func LogoutHandler(c *gin.Context) {
	helpers.DeleteCookie(c)
	c.Redirect(http.StatusFound, "/login")
}
