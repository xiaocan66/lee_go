package binding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	err := Json.BindBody([]byte("{\"foo\":\"pig\"}"), &s)
	assert.NoError(t, err)
	t.Logf("%+v", s)
}
func TestJsonBindingBindMap(t *testing.T) {
	s := make(map[string]string)
	err := Json.BindBody([]byte("{\"name\":\"laowang \"}"), &s)
	assert.NoError(t, err)
	t.Logf("%+v", s)
}
