package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/disaster-information-api/models"
)

func High(c *gin.Context) {
	jsonStr, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
	c.JSON(http.StatusOK, jsonStr)
}
