package binding

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFile struct {
	Filedname string
	Filename  string
	Content   []byte
}

func createRequestMultipartFiles(t *testing.T, file ...testFile) *http.Request {
	var body bytes.Buffer
	mv := multipart.NewWriter(&body)
	for _, f := range file {
		fw, err := mv.CreateFormFile(f.Filedname, f.Filename)
		assert.NoError(t, err)

		n, err := fw.Write(f.Content)
		assert.NoError(t, err)
		assert.Equal(t, len(f.Content), n)

	}
	err := mv.Close()
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/", &body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+";boundary="+mv.Boundary())
	return req
}

func TestFormMultipartBindOneFile(t *testing.T) {
	var s struct {
		FileValue   multipart.FileHeader     `form:"file"`
		FilePtr     *multipart.FileHeader    `form:"file"`
		SliceValues []multipart.FileHeader   `form:"file"`
		SlicePtrs   []multipart.FileHeader   `form:"file"`
		ArrayValues [1]multipart.FileHeader  `form:"file"`
		ArrayPtrs   [1]*multipart.FileHeader `form:"file"`
	}
	file := testFile{"file", "file1", []byte("hell oworld")}
	req := createRequestMultipartFiles(t, file)
	err := FormMultipart.Bind(req, &s)

	assert.NoError(t, err)
	assertMultipartFileHeader(t, &s.FileValue, file)
	assertMultipartFileHeader(t, s.FilePtr, file)
	assert.Len(t, s.SliceValues, 1)
	assertMultipartFileHeader(t, &s.SliceValues[0], file)
	assert.Len(t, s.SliceValues, 1)
	assertMultipartFileHeader(t, s.ArrayPtrs[0], file)
	assertMultipartFileHeader(t, &s.ArrayValues[0], file)
	assertMultipartFileHeader(t, s.ArrayPtrs[0], file)
	t.Logf("%+v", s)

}

func TestFormMultipartBindingTwoFile(t *testing.T) {
	var s struct {
		SliceValues []multipart.FileHeader  `form:"file"`
		SlicePtr    []*multipart.FileHeader `form:"file"`
		ArrayValues []multipart.FileHeader  `form:"file"`
		ArrayPtr    []*multipart.FileHeader `form:"file"`
	}
	files := []testFile{
		{
			"file", "file1", []byte("i am file1"),
		},
		{
			"file", "file2", []byte("i am file2"),
		},
	}
	req := createRequestMultipartFiles(t, files...)
	err := FormMultipart.Bind(req, &s)
	assert.NoError(t, err)
	assert.Len(t, s.SliceValues, len(files))
	assert.Len(t, s.SlicePtr, len(files))
	assert.Len(t, s.ArrayPtr, len(files))
	assert.Len(t, s.ArrayValues, len(files))
	for i := range files {
		assertMultipartFileHeader(t, &s.SliceValues[i], files[i])
		assertMultipartFileHeader(t, s.SlicePtr[i], files[i])
		assertMultipartFileHeader(t, s.ArrayPtr[i], files[i])
		assertMultipartFileHeader(t, &s.ArrayValues[i], files[i])
	}
	t.Logf("%+v", s)

}
func TestFormMultipartBindingBindError(t *testing.T) {
	files := []testFile{
		{"file", "file1", []byte("hello i am file1")},
		{"file", "fiel2", []byte("hello  i am file2")},
	}
	for _, tt := range []struct {
		name string
		s    any
	}{
		{
			"wrong type", &struct {
				Files int `form:"file"`
			}{},
		},
		{
			"wrong array size", &struct {
				Files [1]*multipart.FileHeader `form:"file"`
			}{},
		},
		{
			"wrong slice type", &struct {
				Files []int `form:"file"`
			}{},
		},
	} {
		req := createRequestMultipartFiles(t, files...)
		err := FormMultipart.Bind(req, tt.s)
		if err != nil {
			t.Logf("%+v\n", err)
		

		}
			t.Logf("%+v", tt)
	}

}
func assertMultipartFileHeader(t *testing.T, fh *multipart.FileHeader, file testFile) {
	assert.Equal(t, fh.Filename, file.Filename)
	assert.Equal(t, int64(len(file.Content)), fh.Size)
	mf, err := fh.Open()
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(mf)
	assert.NoError(t, err)
	assert.Equal(t, string(body), string(file.Content))
	err = mf.Close()
	assert.NoError(t, err)

}
func TestReflectDemo(t *testing.T) {
	var num int = 1000
	value := reflect.ValueOf(num)
	switch value.Interface().(type) {
	case int:
		t.Log(value.Interface().(int))

	}

}