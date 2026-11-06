package main

import (
	"net/http"

	"github.com/happyhzr/happyweb"
)

func main() {
	r := happyweb.New()
	r.GET("/", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *happyweb.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *happyweb.Context) {
		c.JSON(http.StatusOK, happyweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
