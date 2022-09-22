package migration

import (
	classModel "gp3/features/class/data"
	menteeModel "gp3/features/mentee/data"
	userModel "gp3/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&classModel.Class{})
	db.AutoMigrate(&menteeModel.Mentee{})
}
