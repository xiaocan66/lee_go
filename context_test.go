package lee

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSetValue(t *testing.T) {
	r := Default()
	r.Get("/api/test", func(ctx *Context) {
		p := ctx.Query("p")
		ctx.Set(p, "hello world")

	})

	r.Run(":9000")

}

type User struct {
	Username string    `form:"username"`
	Password string    `form:"password"`
	Update   time.Time `form:"update" time_format:"2006-01-02 15:04:05"`
	Age      uint8     `form:"age"`
}

func test(ctx *Context) {
	var user User

	log.Println("start...")
	if err := ctx.ShouldBind(&user); err != nil {
		fmt.Fprintf(ctx.Writer, "%v", err)
		return
	}
	fmt.Fprintf(ctx.Writer, "%+v", user)
}
func test2(c *Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Fprintf(c.Writer, "%+v", err)
		return
	}
	fmt.Fprintf(c.Writer, "%+v", user)

}
func TestShouldBind(t *testing.T) {
	router := Default()
	router.Get("/test", test)
	router.Post("/test", test2)
	router.Run(":9000")

}
