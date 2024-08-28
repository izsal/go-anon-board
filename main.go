package main

import (
	"fmt"

	"github.com/izsal/go-anon-board/config"
	"github.com/izsal/go-anon-board/database"
	"github.com/izsal/go-anon-board/routes"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	fmt.Println(conn)
	r.Serve()
}
