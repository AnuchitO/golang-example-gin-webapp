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

	v1 := router.Group("/api/v1/:key")
	{
			v1.GET("/todoService", func(c *gin.Context) {
					service.TodoService()
					c.String(200, "invork todoService by resource")
				})

			v1.GET("/pathParams/:name", func(c *gin.Context) {
						fmt.Println("GET /:name ")
						name := c.Params.ByName("name")
						key := c.Params.ByName("key")

						c.String(200, "hello name:"+name+"  key:"+key)
					})

			v1.GET("/queryParams", func(c *gin.Context) {
						fmt.Println("GET /queryParams")
						token := c.Request.URL.Query().Get("token")
						c.String(200, "hello "+token)
					})

			v1.GET("/json", func(c *gin.Context) {
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
			v1.POST("/pass/json", func(c *gin.Context) {
						var json UserJSON

						if c.EnsureBody(&json) {
							c.JSON(200, json)
						}
					})
	}




	router.Run(":8080")

}
