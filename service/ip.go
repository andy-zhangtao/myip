package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MyIP struct {
}

type Response struct {
	IP string `json:"ip"`
}

func (ip *MyIP) Get(c *gin.Context) {
	format := c.Query("format")
	if strings.ToLower(format) == "json" {
		c.JSON(http.StatusOK, Response{IP: c.ClientIP()})
		return
	}

	c.String(http.StatusOK, "%s", c.ClientIP())
}
