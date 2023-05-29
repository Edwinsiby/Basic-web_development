package main

import (
	"fmt"
	"main/DB"
	"main/Handlers"
	"main/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Can't load env")
	}
	router := gin.New()
	DB.Db, _ = gorm.Open(postgres.Open(os.Getenv("DBS")), &gorm.Config{})
	DB.Db.AutoMigrate(&models.User{})
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", Handlers.IndexHandler)
	router.GET("/signup", Handlers.SignupHandler)
	router.POST("/signuppost", Handlers.SignupPost)
	router.GET("/login", Handlers.LoginHandler)
	router.POST("/loginpost", Handlers.LoginPost)
	router.GET("/home", Handlers.HomeHandler)
	router.GET("/logout", Handlers.LogoutHandler)

	router.GET("/admin", Handlers.AdminHandler)
	router.GET("/admin/edit", Handlers.EditHandler)
	router.GET("/admin/delete", Handlers.DeleteHandler)
	router.POST("/update", Handlers.UpdateHandler)
	router.GET("/loadcreate", Handlers.LoadcreateHandler)
	router.POST("/create", Handlers.CreateHandler)
	router.POST("/search", Handlers.SearchHandler)

	router.Run(":8080")
}
