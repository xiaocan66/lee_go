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
	// key := req.Method + "_" + req.URL.Path
	// method := engine.router[key]
	// if method == nil {
	// 	fmt.Fprintf(w, "Not Found :%s", req.URL.Path)
	// } else {
	// 	method(w, req)
	// }

	c := newContext(w, req)
	engine.router.handle(c)
}


func (engine *Engine) AddRoute(method string, pattern string, hander HandlerFunc) {
	key := method + "_" + pattern
	engine.router.addRouter(key, pattern, hander)
}
// Get defines the method to add Get request
func (engine *Engine) Get(pattern string, hander HandlerFunc) {
	engine.AddRoute("GET", pattern, hander)
}

//Post defines the method to add Post request
func (engine *Engine) Post(pattern string, hander HandlerFunc) {
	engine.AddRoute("POST", pattern, hander)
}

func (engine *Engine) Delete(pattern string, hander HandlerFunc) {
	engine.AddRoute("Delete", pattern, hander)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
