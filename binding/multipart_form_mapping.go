package binding

import (
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
)

type multipartRequest http.Request

var _ setter = (*multipartRequest)(nil)

var (
	ErrMultipartHeader           = errors.New("unsupported field type form multipart.FileHeader")
	ErrMultipartHeaderLenInvalid = errors.New("unsupport len of array for []*multipart.FileHeader ")
)

func (r *multipartRequest) TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (bool, error) {
	if files := r.MultipartForm.File[key]; len(files) != 0 {
		return setByMutipartForm(value, field, files)
	}
	return setByForm(value, field, r.MultipartForm.Value, key, opt)
}

func setByMutipartForm(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		switch value.Interface().(type) {
		case *multipart.FileHeader:
			value.Set(reflect.ValueOf(files[0]))
			isSet, err = true, nil
			return
		}
	case reflect.Struct:
		switch value.Interface().(type) {
		case multipart.FileHeader:
			value.Set(reflect.ValueOf(*files[0]))
			isSet, err = true, nil
			return
		}
	case reflect.Slice:
		slice := reflect.MakeSlice(value.Type(), len(files), len(files))
		isSet, err = setArrayOfMultipartFormFiles(slice, field, files)
		if err != nil && !isSet {
			return
		}
		value.Set(slice)
		return
	case reflect.Array:
		return setArrayOfMultipartFormFiles(value, field, files)
	}

	return false, ErrMultipartHeader
}

// setArrayOfMultipartFormFile
func setArrayOfMultipartFormFiles(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {

	if value.Len() != len(files) {
		return false, ErrMultipartHeaderLenInvalid
	}
	for i := range files {
		set, err := setByMutipartForm(value.Index(i), field, files[i:i+1])
		if !set && err != nil {
			return set, nil
		}

	}

	return true, nil
}

