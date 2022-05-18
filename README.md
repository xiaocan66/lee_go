# Lee_go

## 从零实现一个简易Web框架



## 开启一个HTTP服务

```go
package main
import "lee"

func main(){
    r := lee.Default()
    r.Get("/",func(ctx *lee.Context){
        fmt.Fprint(ctx.Writer,"Hello World")    
    })
    r.Run(":9000")    
}
```
