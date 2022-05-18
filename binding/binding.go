package binding

import (
	"net/http"
)

const (
	MIMEJSON = "application/json"
)

var (
	Form = formBiding{}
)

type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

// Default 根据请求方法和Content-Type 返回合适的Binding实例
func Default(method, contentType string) Binding {
	if method == http.MethodGet {
		return Form
	}
	return nil
}
