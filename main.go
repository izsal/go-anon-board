package main

import (
	"github.com/izsal/go-anon-board/config"
	"github.com/izsal/go-anon-board/controllers"
	"github.com/izsal/go-anon-board/database"
	"github.com/izsal/go-anon-board/routes"
	"github.com/izsal/go-anon-board/services"
)

func main() {
	//config
	c := config.NewConfig()

	//database
	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	// services
	ts := services.NewThreadService(conn)
	rs := services.NewReplyService(conn)

	//controllers
	tc := controllers.NewThreadController(ts)
	rc := controllers.NewReplyController(rs)
	ac := controllers.NewAdminController(ts, rs)

	//routes
	r.RegisterThreadRoutes(tc)
	r.RegisterReplyRoutes(rc)
	r.RegisterAdminRoutes(ac)

	//start server
	r.Serve()
}
