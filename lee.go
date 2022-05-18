package lee

import (
	"net/http"
	"path"
	"text/template"
)

// HandlerFunc defines the request handler use by lee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router       *router
	groups       []*RouterGroup     // store all groups
	htmlTemplate *template.Template //for html render
	funcMap      template.FuncMap   // for html render
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine //all groups share a Engine instance
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine

}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	part := parsePattern(req.URL.Path)
	var prefix string = "/"
	if len(part) > 0 {
		prefix += part[0]
	}
	for _, group := range engine.groups {
		if group.prefix == "" || prefix == group.prefix {
			middlewares = append(middlewares, group.middlewares...)
		}

	}
	c := newContext(w, req)
	c.engine = engine
	c.handlers = middlewares
	engine.router.handle(c)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
func (engine *Engine) LoadHtmlGlob(pattern string) {
	engine.htmlTemplate = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))

}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)

}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, hander HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, hander)

}

// Get defines the method to add Get request
func (group *RouterGroup) Get(pattern string, hander HandlerFunc) {

	group.addRoute("GET", pattern, hander)
}

//Post defines the method to add Post request
func (group *RouterGroup) Post(pattern string, hander HandlerFunc) {
	group.addRoute("POST", pattern, hander)
}

// Delete defines the method to add Delete request
func (group *RouterGroup) Delete(pattern string, hander HandlerFunc) {
	group.addRoute("Delete", pattern, hander)
}

//Run defines the method to start a http server
func (group *RouterGroup) Run(addr string) error {
	return http.ListenAndServe(addr, group.engine)
}

// create static handler

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileserver := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		file := ctx.Param("filepath")
		// check  if file exist  and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		fileserver.ServeHTTP(ctx.Writer, ctx.Req)

	}

}
func (group *RouterGroup) Static(relativePath, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := relativePath + "/*filepath"
	//register GET handler
	group.Get(urlPattern, handler)
}
