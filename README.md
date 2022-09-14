# Lee_go

## 从零实现一个简易Web框架

## 目前已实现的功能

- 动态路由

- 路由分组

- 添加中间件

- HTML模板

- 数据映射 开发中.....

## 功能演示

### 开启一个HTTP服务

```go
package main
import "lee"

func main(){
      r := lee.Default()
    r.Get("/", func(ctx *lee.Context) {
        fmt.Fprint(ctx.Writer, "Hello world")
    })
    r.Run(":9000")  
}
```

### 使用中间件

```go
package main
import "lee"

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

func main(){
      r := lee.Default()
    r.Use(auth()) // 使用中间件
    r.Get("/", func(ctx *lee.Context) {
        fmt.Fprint(ctx.Writer, "Hello world")
    })
    r.Run(":9000")  
}
```

### 动态路由

```go
func main(){
      r := lee.Default()
    r.Use(auth()) // 使用中间件
    r.Get("/user/:username", func(ctx *lee.Context) {
        fmt.Fprint(ctx.Writer, "Hello world")
    })
    r.Get("/assets/*filepath", func(ctx *lee.Context) {
        fmt.Fprint(ctx.Writer, "Hello world")
    })
    r.Run(":9000")  
}
```

### 路由分组

```go
package main
import "lee"

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


func main(){
      r := lee.Default()
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
```
