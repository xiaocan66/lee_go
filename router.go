package lee

import "log"

type router struct {
	handers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method string, pattern string, hander HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "_" + pattern
	r.handers[key] = hander
}
func (r *router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handers[key]; ok {
		handler(c)

	}

}
