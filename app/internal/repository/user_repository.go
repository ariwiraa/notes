package repository

import (
	"errors"
	"notes/app/database"
	"notes/app/entity"

	"github.com/labstack/echo/v4"
)

//UserRepository connects entity.User with Database
// type UserRepository struct {
// 	db *gorm.DB
// }

// func NewUserRepository(db *gorm.DB) *UserRepository {
// 	return &UserRepository {
// 		db: db,
// 	}
// }

func InsertUser(c echo.Context, u *entity.User) (entity.User, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.User{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.User{}, errors.New("Failed to add user")
	}

	return *u, nil
}

func FindAllUsers(c echo.Context) ([]entity.User, error) {
	var user []entity.User
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&user)
	if err.Error != nil {
		return user, err.Error
	}

	return user, nil
}

func FindUserById(c echo.Context, id uint32) (entity.User, error) {
	var u entity.User
	db := database.GetDB(c)

	err := db.Debug().First(&u, "id = ?", id)
	if err.Error != nil {
		return entity.User{}, errors.New("User not Found")
	}

	return u, nil
}

func UpdateUser(c echo.Context, id string, user *entity.User) (int64, error) {
	db := database.GetDB(c)

	err := db.Debug().Model(&entity.User{}).Where("id = ?", id).Updates(user)

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func DeleteUser(c echo.Context, id string) (int64, error) {
	db := database.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.User{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
