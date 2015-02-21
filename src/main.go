package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"service"
)

func rootPath(c *gin.Context) {
	fmt.Println("======= rootPath =====")
	fmt.Println("GET / ")
	unixtime := int32(time.Now().Unix())
	fmt.Println(unixtime)
	c.String(200, "hello world")
}

func main() {

	router := gin.Default()
	router.GET("/", rootPath)

	router.GET("/api/v1/:key/todoService", func(c *gin.Context) {
		service.TodoService()
		c.String(200, "invork todoService by resource")
	})

	router.GET("/api/v1/:key/pathPrarams/:name", func(c *gin.Context) {
		fmt.Println("GET /:name ")
		name := c.Params.ByName("name")

		c.String(200, "hello name:"+name)
	})

	router.GET("/api/v1/:key/queryParams", func(c *gin.Context) {
		fmt.Println("GET /api/v1/:key/queryParams")
		token := c.Request.URL.Query().Get("token")
		c.String(200, "hello "+token)
	})

	router.GET("/api/vi/:key/json", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name = "Nong"
		msg.Message = "hello Json"
		msg.Number = 123456789

		c.JSON(200, msg)
	})

	type UserJSON struct {
		UserName string `json:"user" binding:"required"`
		Phone    string
	}

	/*{
	  	"user":"Anuchit",
	 	"phone":"1234567897"

	}
	*/
	router.POST("/api/v1/:key/pass/json", func(c *gin.Context) {
		var json UserJSON

		if c.EnsureBody(&json) {
			c.JSON(200, json)
		}
	})

	router.Run(":8080")

}
