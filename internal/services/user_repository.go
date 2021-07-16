package services

import (
	"GoLandPruebas/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

var (
	UserService = &userService{setupDB()}
)

func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	return db
}

func (u *userService)GetUserByName(name string, password string) (models.User, error) {
	var user models.User
	u.db.Model(&models.User{}).Where("name = ?", name).Where("password = ?", password).First(&user)
	return user, u.db.Error
}

func (u *userService)GetAllUsers() ([]models.User, error) {
	var users []models.User
	u.db.Model(&models.User{}).First(&users)
	return users, u.db.Error
}

func (u *userService)GetUserById(id uint) (models.User, error) {
	var user models.User
	u.db.Model(&models.User{}).Where("id = ?", id).First(&user)
	return user, u.db.Error
}

func (u *userService)CreateUser(name string, password string) error {
	user := models.User{
		Name: name,
		Password: password,
	}

	u.db.Model(&models.User{}).Create(&user)
	return u.db.Error
}