package request

import (
	"log"
	"math"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	"go-backend/dao"
)

func Employee(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	requests, err := dao.GetRequestsForEmployee(uint(userId), limit, page)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get requests",
		})
		return
	}

	count, err := dao.GetRequestsCountForEmployee(uint(userId))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get requests",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requests":    requests,
		"total_pages": int(math.Max(1, math.Ceil(float64(count)/float64(limit)))),
	})
}
