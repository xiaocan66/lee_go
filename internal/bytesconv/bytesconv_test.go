package bytesconv

import (
	"testing"
)

func TestStringToBytes(t *testing.T) {
	var str = "hello world"
	bts := StringToBytes(str)
	t.Log(string(bts))
}
func TestBytesToString(t *testing.T) {
	var bt = []byte{'h', 'e', 'l', 'l', ' ', 'w', 'o', 'r','l', 'd'}
	str := BytesToString(bt)
	t.Log(str)

}
