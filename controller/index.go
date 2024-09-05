package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hubhike/go-fly/models"
)

func Index(c *gin.Context) {
	jump := models.FindConfig("JumpLang")
	if jump != "cn" {
		jump = "en"
	}
	c.Redirect(302, "/index_"+jump)
}
