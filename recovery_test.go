package lee

import (
	"fmt"
	"testing"
)

func TestRecovery(t *testing.T) {
	engine := Default()
	engine.Get("/panic", func(ctx *Context) {
		var arr = []int{1, 3}
		fmt.Println(arr[3])
	})
	engine.Get("/hello", func(ctx *Context) {

		fmt.Fprint(ctx.Writer, "我还活着！！！")

	})
	if err := engine.Run(":9000"); err != nil {
		t.Log(err)
	}

}
