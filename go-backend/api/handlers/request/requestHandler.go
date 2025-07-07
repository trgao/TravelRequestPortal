package request

import (
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"

	"go-backend/dao"
)

type NewRequest struct {
	StartLocation string `json:"start_location"`
	Destination   string `json:"destination"`
	DateTime      string `json:"date_time"`
	Purpose       string `json:"purpose"`
}

type RequestBody struct {
	Request NewRequest `json:"request"`
}

func Request(c *gin.Context) {
	var body RequestBody

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is not of correct format",
		})
		return
	}

	userIdValue, ok := c.Get("UserID")
	if !ok {
		log.Println("Unable to get user id")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	userId, ok := userIdValue.(float64)
	if !ok {
		log.Println("Invalid user ID type")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user id format",
		})
		return
	}

	dateTime, err := time.Parse("2006-01-02T15:04", body.Request.DateTime)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Date time must be in format YYYY-MM-DDTHH:MM",
		})
		return
	}

	err = dao.CreateRequest(body.Request.StartLocation, body.Request.Destination, dateTime.Add(time.Hour*-8), body.Request.Purpose, uint(userId))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully submitted travel request",
	})
}
