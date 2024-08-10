package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/wendellnd/sqlc/internal/db"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// Create a category
	createCategoryParams := db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Programming",
		Description: "Programming courses",
	}

	err = queries.CreateCategory(ctx, createCategoryParams)
	if err != nil {
		panic(err)
	}

	// Get a category
	category, err := queries.GetCategory(ctx, createCategoryParams.ID)
	if err != nil {
		panic(err)
	}

	println(category.ID, category.Name, category.Description)

	// Update category
	updateCategoryParams := db.UpdateCategoryParams{
		ID:          createCategoryParams.ID,
		Name:        "Programming",
		Description: "Programming courses updated",
	}

	err = queries.UpdateCategory(ctx, updateCategoryParams)
	if err != nil {
		panic(err)
	}

	// List categories
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description)
	}

	// Delete a category
	err = queries.DeleteCategory(ctx, createCategoryParams.ID)
	if err != nil {
		panic(err)
	}

}
