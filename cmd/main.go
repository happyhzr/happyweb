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

	r.GET("/hello", func(c *happyweb.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *happyweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *happyweb.Context) {
		c.JSON(http.StatusOK, happyweb.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
