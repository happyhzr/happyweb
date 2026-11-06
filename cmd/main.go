package main

import (
	"net/http"

	"github.com/happyhzr/happyweb"
)

func main() {
	r := happyweb.New()

	r.GET("/", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello happyweb</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *happyweb.Context) {
			c.HTML(http.StatusOK, "<h1>Hello happyweb</h1>")
		})
		v1.GET("/hello", func(c *happyweb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *happyweb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *happyweb.Context) {
			c.JSON(http.StatusOK, happyweb.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
