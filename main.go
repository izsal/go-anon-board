package main

import (
	"github.com/izsal/go-anon-board/config"
	"github.com/izsal/go-anon-board/routes"
)

func main() {
	c := config.NewConfig()
	r := routes.NewRouter(c)

	r.Serve()
}
