# Lee_go

## 从零实现一个简易Web框架

## 序言

很多时候当我们需要实现一个web应用，第一时间想到的就是去使用哪个框架,然而不同的框架有不同的设计理念,提供的功能也有很大的差别  ,比如Java的`Spring` ,Python的`flask` 、`django` ,Go的 `Beego`、`Gin` 、`Iris` 等。那我们为什么不直接使用语言提供的标准库编写呢？要回答这个问题，首先明白框架的核心应该为我们提供什么？ 为什么要用框架？ 只有理解这些才能知道我们在框架中需要实现那些功能。为了深入理解Gin框架的代码和设计。本项目将会参考**Gin**框架实现Gin框架中部分的功能,学习一门技术最好的方式就是看懂后自己去实现一遍。**Gin**框架的代码一共约**1万4千行**，其中测试代码**9千行**，也就是说实际代码只有**5千行**, 小而美。非常值得初学者去学习。

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
