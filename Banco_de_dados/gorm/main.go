package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    1,
		Name:  name,
		Price: price,
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// create
	//db.Create(&Product{
	//	Name:  "Mouse",
	//	Price: 1000.00,
	//})

	//products := []Product{
	//	{Name: "Notebook", Price: 1000.00},
	//	{Name: "Notebook", Price: 100.00},
	//	{Name: "Notebook", Price: 10.00},
	//}
	//db.Create(products)

	// select one
	var product Product
	//db.First(&product, 1)
	//fmt.Println(product)

	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)

	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	/*
		var products []Product
		db.Limit(2).Offset(2).Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}

		db.Where("price > ?", 100).Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}

		db.Where("name LIKE ?", "%ouse%").Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}
	*/

	var p Product
	db.First(&p, 1)

	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)

	fmt.Println(p2.Name)
	db.Delete(&p2)
}
