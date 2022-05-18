package lee

import (
	"log"
	"net/http"
	"time"
)

type consoleColorModoValue int

const (
	autoColor consoleColorModoValue = iota
	disableColor
	forceColor
)

var consoleColorMode = autoColor

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

type LogFormatter func(params LogFormatterParams) string
type LoggerConfig struct {
	Formatter LogFormatter
}

type LogFormatterParams struct {
	Request *http.Request
	// TimeStamp 显示服务器响应请求的时间
	TimeStamp time.Time
	// StatusCode http响应码
	StatusCode int
	//Latency 请求所花费的时间
	Latency time.Duration
	// ClientIP 显示客户端请求IP
	ClientIP string
	// Path 客户端请求路径
	Path string
	//Method 请求方法
	Method string
	// BodySize 响应体大小
	BodySize int
	// isTerm 输出是否指向控制台
	isTerm bool
	// ErrMessage  请求过程中出现错误.
	ErrMessage string
	//Keys 请求上下文中设置的键
	Keys map[string]any
}

func (formatter *LogFormatterParams) StatusCodeColor() string {
	code := formatter.StatusCode
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red

	}
}
func (formatter *LogFormatterParams) MethodColor() string {
	method := formatter.Method
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset

	}

}
func (formatter *LogFormatterParams) Reset() string {
	return reset
}

func Logger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, t)
	}
}
