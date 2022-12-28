package handler

import (
	"net/http"
	jwt "work_order/pkg/jwtauth"
	"work_order/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NoFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	logger.Infof("NoRoute claims: %#v\n", claims)
	c.JSON(http.StatusOK, gin.H{
		"code":    "404",
		"message": "not found",
	})
}
