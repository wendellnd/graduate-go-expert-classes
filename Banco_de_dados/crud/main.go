package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)

	err = insertProduct(db, *product)
	if err != nil {
		panic(err)
	}

	product.Price = 100.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	product2, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(product2)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	err = deleteProduct(db, product2.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	return err
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	//err = stmt.QueryRowContext(context.Background(), id).Scan(&p.ID, &p.Name, &p.Price)
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
