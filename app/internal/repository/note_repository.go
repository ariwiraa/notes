package repository

import (
	"errors"
	"notes/app/database"
	"notes/app/entity"

	"github.com/labstack/echo/v4"
)

func InsertNote(c echo.Context, n *entity.Note) (entity.Note, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&n)
	if err != nil {
		return entity.Note{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.Note{}, errors.New("Failed to add user")
	}

	return *n, nil
}

func FindAllNotes(c echo.Context) ([]entity.Note, error) {
	var note []entity.Note
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&note)
	if err.Error != nil {
		return note, err.Error
	}

	return note, nil
}

func FindNoteById(c echo.Context, id uint32) (entity.Note, error) {
	var u entity.Note
	db := database.GetDB(c)

	err := db.Debug().First(&u, "id = ?", id)
	if err.Error != nil {
		return entity.Note{}, errors.New("User not Found")
	}

	return u, nil
}

func UpdateNote(c echo.Context, id string, note *entity.Note) (int64, error) {
	db := database.GetDB(c)

	err := db.Debug().Model(&entity.Note{}).Where("id = ?", id).Updates(note)

	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func DeleteNote(c echo.Context, id string) (int64, error) {
	db := database.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Note{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
