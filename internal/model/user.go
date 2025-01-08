package model

type User struct{
	UserId 	 uint 	`gorm:"primaryKey;autoIncrement"`
	Usermame string `gorm:"uniqueIndex"`
	Password string
}

