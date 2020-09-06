package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {

	ids := IDService{}
	ids.Init(1)


	r := gin.Default()
	r.GET("/generate", func(c *gin.Context) {

		str := ""
		id, err := ids.GenerateNewID()
		if err != nil {
			str = "there is an error"
		}else {
			str = strconv.FormatInt(id,10)
		}
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}