package main

import (
	"fmt"
	"log"
	"project/route"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"

	"project/orm/model"
)

type Querier interface {
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := "host=db user=user password=postgres dbname=crud port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Article{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migration: ok")

	g.UseDB(db)

	router := route.SetupRouter()
	router = route.PostUser(router)

	router.Run()
}
