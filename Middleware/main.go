package main

import (
	"io"
	"log"
	"main/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "world"})
}

func postmethod(ctx *gin.Context) {
	body := ctx.Request.Body
	bodys, _ := io.ReadAll(body)
	ctx.JSON(http.StatusOK, gin.H{"message": "Do something",
		"body": string(bodys),
	})

}

// http://localhost:9090/hellogetquery?name=farhan&age=27
func hellogetquery(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	age, ageok := ctx.GetQuery("age")
	if !ageok {
		log.Fatalf("No query param provided")
	}

	if !ok {
		log.Fatalf("No query param provided")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": "http.StatusOK",
		"is_success":  "true",
		"data":        "In hellogetquery",
		"message":     "Success",
		"name":        name,
		"age":         age,
	})

}

// http://localhost:9090/hellogeturldata/farhan/52
func hellogeturldata(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": "http.StatusOK",
		"is_success":  "true",
		"data":        "nil",
		"message":     "Success",
		"name":        name,
		"age":         age,
	})

}

func main() {
	router := gin.Default()

	// So this is basic auth system for API's
	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})

	admin := router.Group("/admin", auth)
	{
		admin.GET("/hello", hello)
	}
	//router.GET("/hello", hello)
	router.GET("/hellopost", middleware.Authenticate, postmethod)
	router.GET("/hellogetquery", hellogetquery)
	router.GET("/hellogeturldata/:name/:age", hellogeturldata)
	router.Run(":9090")
}
