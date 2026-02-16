package demo

import (
	"chat-system/config"
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}




func gormDemo() {
	db, err := gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42
	fmt.Printf("value is %v\n", products)
	// Update - update product's price to 200
	rowsAffected, err := gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	fmt.Printf("affetRow is %v\n", rowsAffected)
	// Update - update multiple fields
	rowsAffected, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})
	fmt.Printf("affetRow is %v\n", rowsAffected)
	// Delete - delete product
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}
