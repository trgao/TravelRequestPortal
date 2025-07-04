package request

import (
	"log"

	"github.com/gin-gonic/gin"
	"net/http"

	"go-backend/dao"
)

func Summary(c *gin.Context) {
	userId, ok := c.Get("UserID")
	if !ok {
		log.Println("Unable to get user id")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	pending, err := dao.GetPending(uint(userId.(float64)))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get pending count",
		})
		return
	}

	approved, err := dao.GetApproved(uint(userId.(float64)))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get approved count",
		})
		return
	}

	rejected, err := dao.GetRejected(uint(userId.(float64)))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get rejected count",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pending":  pending,
		"approved": approved,
		"rejected": rejected,
	})
}
