package DB

import (
	"main/models"

	"gorm.io/gorm"
)

var Db *gorm.DB
var UserList []models.User
