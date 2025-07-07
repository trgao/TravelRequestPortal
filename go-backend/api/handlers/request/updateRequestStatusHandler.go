package request

import (
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	"go-backend/dao"
)

type UpdateRequestStatusBody struct {
	Status string `json:"status" binding:"required,oneof=pending approved rejected"`
}

func UpdateRequestStatus(c *gin.Context) {
	var request UpdateRequestStatusBody

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is not of correct format",
		})
		return
	}

	requestId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request id format",
		})
		return
	}

	err = dao.UpdateRequestStatus(requestId, request.Status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update request status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated travel request status",
	})
}
