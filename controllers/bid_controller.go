package controllers

import (
	"go-api/models"
	"go-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BidHandler(c *gin.Context) {
	var requestPayload utils.BidRequest

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := utils.ValidateBidRequest(requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !models.IsIdPresent(requestPayload.Id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found in Redis"})
		return
	}

	response, contentType, err := models.GetResponseData(requestPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", contentType)
	c.String(http.StatusOK, response)
}
