package main

import (
	"github.com/izsal/go-anon-board/config"
	"github.com/izsal/go-anon-board/controllers"
	"github.com/izsal/go-anon-board/database"
	"github.com/izsal/go-anon-board/routes"
	"github.com/izsal/go-anon-board/services"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	ts := services.NewThreadService(conn)

	tc := controllers.NewThreadController(ts)

	r.RegisterThreadRoutes(tc)

	r.Serve()
}
