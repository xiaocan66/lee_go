package binding

import (
	"testing"
	"time"
)

func TestSetFormMap(t *testing.T) {
	var res = make(map[string]string)
	var m = make(map[string][]string)
	m["name"] = []string{
		"xiaocan",
	}
	m["age"] = []string{
		"11",
	}
	m["sex"] = []string{
		"nan",
	}
	err := setFormMap(res, m)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(res)
	}
}

func TestMapFormByTag(t *testing.T) {
	var user struct {
		Username string    `tag:"username"`
		Password string    `tag:"password"`
		Email    string    `tag:"email"`
		Array    []int     `tag:"array"`
		UpdateAt time.Time `tag:"update_at" time_format:"2006-01-02 15:04:05" time_location:"Asia/Shanghai"`
	}
	var form = make(map[string][]string)

	form["username"] = []string{"lizican"}
	form["password"] = []string{"12334"}
	form["email"] = []string{"94902@qq.com"}
	form["update_at"] = []string{"2022-03-27 10:21:44"}
	form["array"] = []string{"1", "2", "3", "4", "4", "4", "5", "5"}

	err := mapFormByTag(&user, form, "tag")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("%+v", user)
	}
}
