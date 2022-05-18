package binding

import "net/http"

type formBiding struct{}

func (formBiding) Name() string {
	return "form"
}

func (formBiding) Bind(req *http.Request, obj any) error {
	return nil
}
