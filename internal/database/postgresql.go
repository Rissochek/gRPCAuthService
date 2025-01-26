package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"AuthProject/internal/model"
	"AuthProject/internal/utils"
)

func InitDataBase() (*gorm.DB){
	dsn := "host=localhost user=rissochek password=123 dbname=auth_db port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to init DB: %v", err)
	}
	db.AutoMigrate(&model.User{})
	return db
}

func AddUserToDataBase(db *gorm.DB, user *model.User) error{
	hashed_user_password := utils.GenerateHash(user.Password)
	user.Password = hashed_user_password

	result := db.Create(user)
	if result != nil{
		return result.Error
	}
	return nil
}

func SearchUserInDB(db *gorm.DB, user *model.User) error{
	target_user := model.User{}
	if err := db.Where(&model.User{Usermame: user.Usermame}).First(&target_user).Error; err != nil{
		return err
	}
	err := utils.CompareHashAndPassword(user.Password, target_user.Password)
	return err
}

