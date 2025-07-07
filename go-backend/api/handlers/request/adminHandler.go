package request

import (
	"log"
	"math"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	"go-backend/dao"
)

func Admin(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id format",
		})
		return
	}

	employee := c.Query("employee")
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid limit format",
		})
		return
	}

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page format",
		})
		return
	}

	requests, err := dao.GetRequestsForAdmin(uint(userId), employee, limit, page)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get requests",
		})
		return
	}

	count, err := dao.GetRequestsCountForAdmin(uint(userId), employee)
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
