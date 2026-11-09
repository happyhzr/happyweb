package main

import (
	"log"
	"net/http"
	"time"

	"github.com/happyhzr/happyweb"
)

func main() {
	r := happyweb.New()

	r.GET("/", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello happyweb</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *happyweb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}

func onlyForV2() happyweb.HandlerFunc {
	return func(c *happyweb.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
