package main

import (
	"log"
	"net/http"
	"time"

	"gee"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	// r := gee.New()
	// r.Use(gee.Logger())
	// r.GET("/index", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Index page</h1>")
	// })
	// r.GET("/", func(ctx *gee.Context) {
	// 	ctx.HTML(http.StatusOK, "<h1> haha </h1>")
	// })

	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/", func(c *gee.Context) {
	// 		c.HTML(http.StatusOK, "<h1>hello gee</h1>")
	// 	})
	// 	v1.GET("/hello", func(c *gee.Context) {
	// 		c.String(http.StatusOK, "hello %s, you're at %s \n", c.Query("name"), c.Path)
	// 	})
	// }

	// v2 := r.Group("/v2")
	// v2.Use(onlyForV2())
	// {
	// 	v2.GET("/hello/:name", func(ctx *gee.Context) {
	// 		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	// 	})
	// 	v2.POST("/login", func(c *gee.Context) {
	// 		c.JSON(http.StatusOK, gee.H{
	// 			"username": c.PostForm("username"),
	// 			"password": c.PostForm("password"),
	// 		})
	// 	})
	// }

	// r.Run(":9999")
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		println()
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
