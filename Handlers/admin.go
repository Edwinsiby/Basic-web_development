package Handlers

import (
	"fmt"
	"main/DB"
	"main/helpers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		if err := DB.Db.Find(&DB.UserList).Error; err != nil {
			fmt.Println("user not found")
		}
		c.HTML(http.StatusOK, "adminpanel.html", gin.H{
			"data": DB.UserList,
		})
	}

}

func EditHandler(c *gin.Context) {

	id := c.Query("id")
	if err := DB.Db.Where("id = ?", id).Find(&DB.UserList).Error; err != nil {
		fmt.Println("user not found")
	}
	c.HTML(http.StatusOK, "adminedit.html", gin.H{
		"data": DB.UserList,
	})
}

func DeleteHandler(c *gin.Context) {
	id := c.Query("id")
	var user models.User
	result := DB.Db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		panic("failed to delete user")
	} else {
		c.Redirect(http.StatusFound, "/admin")
	}

}

func UpdateHandler(c *gin.Context) {
	id := c.Query("id")
	var user models.User
	user.Name = c.Request.FormValue("Username")
	user.Email = c.Request.FormValue("Usermail")
	user.Password = c.Request.FormValue("Password")
	user.Role = c.Request.FormValue("Role")
	result := DB.Db.Where("id=?", id).Updates(&user)
	if result.Error != nil {
		panic("failed to update user")
	} else {
		c.Redirect(http.StatusFound, "/admin")
	}
}

func LoadcreateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admincreate.html", nil)
}

func CreateHandler(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("Username")
	user.Email = c.Request.FormValue("Usermail")
	user.Password = c.Request.FormValue("Password")
	user.Role = c.Request.FormValue("Role")
	DB.Db.Save(&user)
	c.Redirect(http.StatusFound, "/admin")
}

func SearchHandler(c *gin.Context) {
	name := c.Request.FormValue("Search")
	if err := DB.Db.Where("name = ?", name).Find(&DB.UserList).Error; err != nil {
		fmt.Println("user not found")
	}
	c.HTML(http.StatusOK, "adminpanel.html", gin.H{
		"data": DB.UserList,
	})
}
