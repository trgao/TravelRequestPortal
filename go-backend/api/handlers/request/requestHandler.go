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
	}

	userId, ok := c.Get("UserID")
	if !ok {
		log.Println("Unable to get user id")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	dateTime, err := time.Parse("2006-01-02T15:04", body.Request.DateTime)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Date time is not the correct format",
		})
		return
	}

	err = dao.CreateRequest(body.Request.StartLocation, body.Request.Destination, dateTime.Add(time.Hour*-8), body.Request.Purpose, uint(userId.(float64)))
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
