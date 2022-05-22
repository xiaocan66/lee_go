package binding

import (
	"net/http"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTL               = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
)

var (
	Form          = formBiding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}

	Json = jsonBinding{}
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
	switch contentType {
	case MIMEJSON:
		return Json
	case MIMEMultipartPOSTForm:
		return FormMultipart
	case MIMEPOSTForm:
		return FormPost
	default:
		return Form

	}

}
