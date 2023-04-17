package handler

import (
	usecase "female-daily/app/usecase"
	"female-daily/config"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Handler interface{}

var once = &sync.Once{}

type handler struct {
	usecase *usecase.Usecase
}

func Init(usecase *usecase.Usecase) Handler {
	var h *handler
	once.Do(func() {
		h = &handler{
			usecase: usecase,
		}
		h.Serve()
	})
	return h
}

func (h *handler) Serve() {
	router := gin.Default()
	group := router.Group("/api/v1")
	group.GET("/fetch", h.FetchUser)
	group.GET("/user/:id", h.GetByID)
	group.GET("/user", h.GetList)
	group.POST("/user", h.CreateUser)
	group.PUT("/user/:id", h.UpdateUser)
	group.DELETE("/user/:id", h.authenticateMiddleware, h.DeleteUser)

	serverString := fmt.Sprintf("%s:%s", config.AppConfig["host"], config.AppConfig["port"])
	router.Run(serverString)
}

func (h *handler) authenticateMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("authorization")

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		c.Abort()
		return
	}

	if tokenString != "3cdcnTiBsl" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		c.Abort()
		return
	}

	c.Next()
}
