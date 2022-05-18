package test

import (
	"fmt"
	"lee"
	"testing"
)

// auth 定义中间件逻辑
func auth() lee.HandlerFunc {
	return func(ctx *lee.Context) {
		username := ctx.Query("username")
		if username == "lee" {
			ctx.Next()
		}
		ctx.Abort()
	}
}
func auth2() lee.HandlerFunc {
	return func(ctx *lee.Context) {
		username := ctx.Query("username")
		if username == "lee2" {
			ctx.Next()
		}
		ctx.Abort()
	}
}

func TestDefault(t *testing.T) {
	r := lee.Default()
	// r.Use(auth())

	r.Get("/", func(ctx *lee.Context) {
		fmt.Fprint(ctx.Writer, "Hello world")
	})

	g1 := r.Group("v1")
	g1.Use(auth())
	g1.Get("/user", func(ctx *lee.Context) {
		fmt.Fprint(ctx.Writer, "g1")

	})
	g2 := r.Group("v2")
	g2.Use(auth2())
	g2.Get("/post", func(ctx *lee.Context) {
		fmt.Fprint(ctx.Writer, "g2")
	})
	r.Run(":9000")
}
