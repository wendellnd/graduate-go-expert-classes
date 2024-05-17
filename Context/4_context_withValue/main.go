package main

import (
	"context"
	"fmt"
)

type Key string

func main() {
	var key Key = "token"
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	var key Key = "token"
	token := ctx.Value(key)
	fmt.Println(token)
}
