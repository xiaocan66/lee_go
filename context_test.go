package lee

import "testing"

func TestSetValue(t *testing.T) {
	r := Default()
	r.Get("/api/test", func(ctx *Context) {
		p := ctx.Query("p")
		ctx.Set(p, "hello world")

	})
	r.Run(":9000")

}
