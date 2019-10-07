package gin

import (
	"encoding/json"
	"go-pra/copys/httprouter-copy/httprouter"
	"html/template"
	"math"
	"net/http"
	"path"
)

const (
	AbortIndex = math.MaxInt8 / 2
)

type (
	HandlerFunc func(*Context)

	H map[string]interface{}

	ErrorMsg struct {
		Message string      `json:"msg"`
		Meta    interface{} `json:"meta"`
	}

	Context struct {
		Req      *http.Request
		Writer   http.ResponseWriter
		Keys     map[string]interface{}
		Errors   []ErrorMsg
		Params   httprouter.Params
		handlers []HandlerFunc
		engine   *Engine
		index    int8
	}

	RouterGroup struct {
		Handlers []HandlerFunc
		prefix   string
		parent   *RouterGroup
		engine   *Engine
	}

	Engine struct {
		*RouterGroup
		handlers404   []HandlerFunc
		router        *httprouter.Router
		HTMLTemplates *template.Template
	}
)

func New() *Engine {
	engine := &Engine{}
	engine.RouterGroup = &RouterGroup{nil, "", nil, engine}
	engine.router = httprouter.New()
	// engine.router.NotFound = engine.handle404
	return engine
}

// func Default() *Engine {
// 	engine := New()
// 	engine.
// }

func (engine *Engine) LoadHTMLTemplates(pattern string) {
	engine.HTMLTemplates = template.Must(template.ParseGlob(pattern))
}

func (engine *Engine) NotFound404(handlers ...HandlerFunc) {
	engine.handlers404 = handlers
}

func (engine *Engine) handle404(w http.ResponseWriter, req *http.Request) {
	handlers := engine.combineHandlers(engine.handlers404)
	c := engine.createContext(w, req, nil, handlers)
	if engine.handlers404 == nil {
		http.NotFound(c.Writer, c.Req)
	} else {
		c.Writer.WriteHeader(404)
	}

	c.Next()
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.ServeHTTP(w, req)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

/******** ROUTES GROUPING *******/

func (group *RouterGroup) createContext(w http.ResponseWriter, req *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Writer:   w,
		Req:      req,
		index:    -1,
		engine:   group.engine,
		Params:   params,
		handlers: handlers,
	}
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middlewares...)
}

func (group *RouterGroup) Group(component string, handlres ...HandlerFunc) *RouterGroup {
	prefix := path.Join(group.prefix, component)
	return &RouterGroup{
		Handlers: group.combineHandlers(handlres),
		parent:   group,
		prefix:   prefix,
		engine:   group.engine,
	}
}

func (group *RouterGroup) Handle(method, p string, handlers []HandlerFunc) {
	p = path.Join(group.prefix, p)
	handlers = group.combineHandlers(handlers)
	group.engine.router.Handle(method, p, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		// TODO lksdjf
		group.createContext(w, req, params, handlers).Next()
	})
}

func (group *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	group.Handle("POST", path, handlers)
}

func (group *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	group.Handle("GET", path, handlers)
}

func (group *RouterGroup) combineHandlers(handlers []HandlerFunc) []HandlerFunc {
	s := len(group.Handlers) + len(handlers)
	h := make([]HandlerFunc, 0, s)
	h = append(h, group.Handlers...)
	h = append(h, handlers...)
	return h
}

/** FLOW AND ERROR MANAGEMENT **/
func (c *Context) Next() {
	c.index++
	s := int8(len(c.handlers))
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Abort(code int) {
	c.Writer.WriteHeader(code)
	c.index = AbortIndex
}

func (c *Context) Fail(code int, err error) {
	c.Error(err, "Operation aborted")
	c.Abort(code)
}

func (c *Context) Error(err error, meta interface{}) {
	c.Errors = append(c.Errors, ErrorMsg{
		Message: err.Error(),
		Meta:    meta,
	})
}

/** ENCOGING MANAGEMENT */
func (c *Context) JSON(code int, obj interface{}) {
	if code >= 0 {
		c.Writer.WriteHeader(code)
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		c.Error(err, obj)
		http.Error(c.Writer, err.Error(), 500)
	}
}
