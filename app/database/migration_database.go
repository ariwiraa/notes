package database

import (
	"notes/app/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migrations(e *echo.Echo, db *gorm.DB) {
	e.Logger.Info("Mulai dengan Migrasi Database")

	err := db.AutoMigrate(entity.User{}, entity.Note{})
	if err != nil {
		e.Logger.Error(err)
	}
}

// func main() {
// 	e := echo.New()
// 	e.Logger.Info("Menginisialisasikan Database")

// 	config := config.GetConfig(e)
// 	e.Logger.Info(config)

// 	dsn := fmt.Sprintf("host=%s port=%s username=%s password=%s name=%s sslmode=disable",
// 		config.Database.Host,
// 		config.Database.Port,
// 		config.Database.Username,
// 		config.Database.Password,
// 		config.Database.Name)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.AutoMigrate(entity.User{}, entity.Note{})
// 	if err != nil {
// 		e.Logger.Error(err)
// 	}

// }
