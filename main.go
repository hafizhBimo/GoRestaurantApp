package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

// func getFoodMenu(c echo.Context) error {

// 	return c.JSON(200, map[string]interface{}{
// 		"message": "data successfully retrieved",
// 		"data":    foodMenu,
// 	})
// }

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "Bakmie",
			Price:     25000,
		},
		{
			Name:      "Mie Goreng",
			OrderCode: "Mie Goreng",
			Price:     15000,
		},
	}

	fmt.Println(foodMenu)

	dbAddress := "host=localhost port=5432 user=postgres password=Tahumerayap21 dbname=GoRestoApp sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	//to migrate the MenuItem struct into table in db
	// db.AutoMigrate(&MenuItem{})

	db.Create(&foodMenu)
}

func main() {
	seedDB()
	//echo is one of the Go framework that help us to create a REST project
	e := echo.New()

	//get food menu
	// e.GET("/menu/food", getFoodMenu)

	//calling server
	e.Logger.Fatal(e.Start(":8000"))
}
