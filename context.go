package lee

import (
	"encoding/json"
	"fmt"
	"lee/binding"
	"math"
	"net/http"
)

type H map[string]interface{}

// abortIndex 表示一个函数函数终止值
const abortIndex int8 = math.MaxInt8 >> 1

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	engine     *Engine
	Path       string
	Params     map[string]string
	Method     string
	StatusCode int
	handlers   []HandlerFunc
	index      int8
	Keys       map[string]any
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}

}

// Next 处理下一个中间件
func (c *Context) Next() {
	c.index++
	s := int8(len(c.handlers))
	for c.index < s {
		c.handlers[c.index](c)
		c.index++
	}

}

// IsAbort 如果上下文被终止 返回true
func (c *Context) IsAbort() bool {
	return c.index >= abortIndex

}

// Abort 终止请求处理链 继续向下传递
// 注意不会终止当前程序
func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) Set(key string, value any) {
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}
	fmt.Println(c.Keys)
	c.Keys[key] = value
}
func (c *Context) Get(key string) (value any, exist bool) {
	if c.Keys == nil {
		return nil, false
	}
	value, exist = c.Keys[key]
	return

}
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
func (c *Context) SetHeader(key, val string) {
	c.Writer.Header().Set(key, val)

}
func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain;charset=utf-8")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)

}
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplate.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
func (c *Context) Fail(code int, message string) {
	c.Status(code)
	c.Writer.Write([]byte(message))
}

func (c *Context) ContentType() string {

	return filterFlags(c.Req.Header.Get("Content-Type"))
}
// ShouldBind checks the Method and Content-Type  to select a binding engine automatically
// Deepending on the "Content-Type" header different binding are used ,for example
// "application/json" --> Json binding
// "application/xml" --> Xml binding
// It parses the request's body as Json if Content-Type == "application/json" using Json or XML as a Json input
// It decodes the json playload into the  struct specified as a  pointer.
// Like c.Bind() but this method does not set the response status code to 400 or abort if  input is not valid.
func (c *Context) ShouldBind(obj any) error {
	b := binding.Default(c.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

// ShouldBindWith
func (c *Context) ShouldBindWith(obj any, b binding.Binding) error {
	return b.Bind(c.Req, obj)
}
