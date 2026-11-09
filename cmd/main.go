package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/happyhzr/happyweb"
)

type Student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := happyweb.New()
	r.Use(happyweb.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &Student{Name: "happyhzr", Age: 18}
	stu2 := &Student{Name: "hzr", Age: 18}
	r.GET("/", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", happyweb.H{
			"title":  "happyweb",
			"stuArr": []*Student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *happyweb.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", happyweb.H{
			"title": "happyweb",
			"now":   time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
