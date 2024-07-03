package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) Post(c *gin.Context) {
	url := c.Request.URL.String()

	switch {
	case strings.HasPrefix(url, "/user/"):
	}
}
