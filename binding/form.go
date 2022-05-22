package binding

import (
	"errors"
	"log"
	"net/http"
)

type formBiding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

const defaultMemory = 32 << 20 // 默认表单最大大小

func (formBiding) Name() string {
	return "form"
}

func (formBiding) Bind(req *http.Request, obj any) error {
	log.Println("form binding...")
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	if err := mapForm(obj, req.Form); err != nil {
		return err
	}
	// todo 验证值
	return nil
}

// formPostBinding

func (formPostBinding) Name() string {
	return "form-urlencoded"
}
func (formPostBinding) Bind(req *http.Request, obj any) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := mapForm(obj, req.Form); err != nil {
		return nil
	}
	// todo 验证值的正确性
	return nil
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}
func (formMultipartBinding) Bind(req *http.Request, obj any) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(req), "form"); err != nil {
		return err
	}
	return nil
}
