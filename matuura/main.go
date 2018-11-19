package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hal-ms/driver/matuura/repo"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/matuura/model"
)

type Test struct {
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
	Slice [][]int   `json:"slice"`
}

// time 2018-08-16T17:16:08.454+09:00

func main() {
	r := gin.Default()
	fmt.Println(time.Now())
	r.GET("/test", func(c *gin.Context) {
		var test Test
		test.Title = "hoge"
		test.Time = time.Now()
		test.Slice = [][]int{{2, 4, 5}, {4, 5, 6}}

		c.JSON(http.StatusOK, test)
	})

	r.POST("/alarm", func(c *gin.Context) {
		var alarm model.Alarm
		err := c.BindJSON(&alarm)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		repo.SetAlarm(alarm)
	})

	r.GET("/alarms", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.GetAllAlarm())
	})
	r.GET("/alarm/:seet_id", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.FindBySeetIDAlarm(c.Param("seet_id")))
	})
	r.Run()
}
