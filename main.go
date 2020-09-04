package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", Get)
	router.POST("/create", Create)
	router.PUT("/:id/", Update)
	router.DELETE("/:id", Delete)

	router.Run(":8080")
}

func Get(c *gin.Context) {
	m := make(map[string]string)
	m["message"] = "get"
	c.JSON(200, m)
}
func Create(c *gin.Context) {
	m := make(map[string]string)
	m["message"] = "create"
	c.JSON(200, m)
}

func Update(c *gin.Context) {
	m := make(map[string]string)
	m["message"] = c.Param("id")
	c.JSON(200, m)
}

func Delete(c *gin.Context) {
	m := make(map[string]string)
	m["message"] = c.Param("id")
	c.JSON(200, m)
}
