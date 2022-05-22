package binding

import (
	"bytes"
	"errors"
	"io"
	"lee/internal/json"
	"net/http"
)

type jsonBinding struct{}

// EnableDecoderUseNumber  is used to called the UseNumber  method on  the JSON
// Decoder instance 。UseNumber causes the Decoder to unmarshal a number into a
// interface{} as a Number instead of as a float32
var EnableDecoderUseNumber = false

var EnableDecoderDisallowUnknownFileds = false

func (jsonBinding) Name() string {
	return "json"
}
func (jsonBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request ")
	}
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj any) error {
	return decodeJSON(bytes.NewReader(body), obj)
}
func decodeJSON(r io.Reader, obj any) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if EnableDecoderDisallowUnknownFileds {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	// todo valiator
	return nil 

}
