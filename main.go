package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"./model"
)

const DatabaseType = "sqlite3"
const DatabasePath = "test.db"

func main()  {
	db, err := gorm.Open(DatabaseType, DatabasePath)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.Product{})

	// Create
	db.Create(&model.Product{Code: "L1212", Price: 1000})

	// Read
	var product model.Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}