package lee

import (
	"net/http"
)

// HandlerFunc defines the request handler use by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}

}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func (engine *Engine) AddRoute(method string, pattern string, hander HandlerFunc) {
	engine.router.addRoute(method, pattern, hander)
}

// Get defines the method to add Get request
func (engine *Engine) Get(pattern string, hander HandlerFunc) {
	engine.AddRoute("GET", pattern, hander)
}

//Post defines the method to add Post request
func (engine *Engine) Post(pattern string, hander HandlerFunc) {
	engine.AddRoute("POST", pattern, hander)
}

// Delete defines the method to add Delete request
func (engine *Engine) Delete(pattern string, hander HandlerFunc) {
	engine.AddRoute("Delete", pattern, hander)
}

//Run defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
