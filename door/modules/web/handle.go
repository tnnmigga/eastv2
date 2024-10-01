package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *web) initHandle() {
	m.GET("/index", index)
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
