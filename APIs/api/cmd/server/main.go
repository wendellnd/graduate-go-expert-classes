package main

import (
	"fmt"

	"github.com/wendellnd/graduate-go-expert-classes/api/configs"
)

func main() {
	cfg, _ := configs.LoadConfig(".")
	fmt.Println(cfg.DBDriver)
}
