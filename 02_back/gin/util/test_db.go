package util

import (
	"project/orm/model"

	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewTestDB(name string) (*gorm.DB, error) {
	dsn := "host=db user=crud_user password=postgres dbname=crud_test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	txdb.Register(
		name,
		"postgres",
		dsn,
	)

	dialector := postgres.New(postgres.Config{
		DriverName: name,
		DSN:        dsn,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Article{},
	)
	if err != nil {
		panic(err)
	}

	return db, nil
}
