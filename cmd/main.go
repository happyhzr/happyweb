package main

import (
	"net/http"

	"github.com/happyhzr/happyweb"
)

func main() {
	r := happyweb.Default()
	r.GET("/", func(c *happyweb.Context) {
		c.String(http.StatusOK, "Hello happyhzr\n")
	})
	r.GET("/panic", func(c *happyweb.Context) {
		names := []string{"happyhzr"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
