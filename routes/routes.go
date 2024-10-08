package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izsal/go-anon-board/controllers"
	"github.com/izsal/go-anon-board/utils"
)

func setupDefaults(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(utils.CreateApiError(http.StatusNotFound, errors.New("no route found")))
		return
	})
}

func (r *router) RegisterThreadRoutes(c controllers.ThreadController) {
	rg := r.Group("/api/threads")
	rg.GET("/", c.ListThreads)
	rg.POST("/", c.CreateThread)
	rg.GET("/:id", c.GetThread)
	rg.PUT("/:id", c.ReportThread)
	rg.DELETE("/:id", c.DeleteThread)
}

func (r *router) RegisterReplyRoutes(c controllers.ReplyController) {
	rg := r.Group("/api/replies/:tid")
	rg.GET("/", c.ListReplies)
	rg.POST("/", c.CreateReply)
	rg.GET("/:id", c.GetReply)
	rg.PUT("/:id", c.ReportReply)
	rg.DELETE("/:id", c.DeleteReply)
}

func (r *router) RegisterAdminRoutes(c controllers.AdminController) {
	apiKey := r.c.Get().GetString("app.api_key")
	rg := r.Group("/api/admin")
	rg.Use(ApiKey("X-API-KEY", apiKey))

	rg.DELETE("/threads/:id", c.DeleteThread)
	rg.DELETE("/replies/:id", c.DeleteReply)
}
