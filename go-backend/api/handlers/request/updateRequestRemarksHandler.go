package request

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"

	"go-backend/dao"
)

type UpdateRequestRemarksBody struct {
	Remarks string `json:"remarks"`
}

func UpdateRequestRemarks(c *gin.Context) {
	var request UpdateRequestRemarksBody

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
	}

	requestId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	err = dao.UpdateRequestRemarks(requestId, request.Remarks)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update request remarks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated travel request remarks",
	})
}
