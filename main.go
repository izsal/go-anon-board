package main

import (
	"fmt"

	"github.com/izsal/go-anon-board/config"
)

func main() {
	c := config.NewConfig()

	fmt.Println(c)
}
