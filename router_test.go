package lee

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/user/Info/test", nil)
	r.addRoute("GET", "/user/:name", nil)
	r.addRoute("GET", "/user/aaa/*filepath", nil)
	// r.addRoute("GET", "/user/hello/123", nil)
	return r

}
func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/user/:name"), []string{"user", ":name"})
	if !ok {
		t.Log("/user/:name parse failed")
	}
	ok = reflect.DeepEqual(parsePattern("/user/*path"), []string{"user", "*path"})
	if !ok {
		t.Log("/user/* parse failed")
	}
	ok = reflect.DeepEqual(parsePattern("/user/xiaocan/jj"), []string{"user", "xiaocan", "jj"})
	if !ok {
		t.Log("/user/xiaocan parse failed")

	}
	t.Log("parse successfully")

}
func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	var url = "/user/Info/test"
	n, _ := r.getRoute("GET", url)
	if n == nil {
		t.Fatalf("nil shouldn't be returned :%s", url)
	}
	if n.pattern != url {
		t.Fatalf("should match:%s", "/user/Info/test")
	}

	url = "/user/hello"
	n, ps := r.getRoute("GET", url)
	if n == nil {
		t.Fatalf("nil shouldn't be returned %s", url)
	}
	if !reflect.DeepEqual(ps, map[string]string{"name": "hello"}) {
		t.Fatal("参数不正确！！")

	} else {
		t.Logf("获取到的参数:%v", ps)

	}
	url = "/user/aaa/test/a.jpg"
	n, ps = r.getRoute("GET", url)
	if n == nil {
		t.Fatalf("nil shouldn't be returned : %s", url)
	}
	if !reflect.DeepEqual(ps, map[string]string{"filepath": "test/a.jpg"}) {
		t.Fatalf("参数不正确!!")
	} else {
		t.Logf("获取到的参数:%v", ps)
	}

}

