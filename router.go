package httpgin

import (
	"github.com/gin-gonic/gin"
	httpmodule "github.com/things-kit/things-kit-http"
)

// ginRouter wraps gin.Engine or gin.RouterGroup to implement the http.Router interface.
// This allows route registration to be framework-agnostic.
type ginRouter struct {
	router gin.IRouter
}

// newGinRouter creates a new wrapper around gin.IRouter (Engine or RouterGroup).
func newGinRouter(r gin.IRouter) httpmodule.Router {
	return &ginRouter{router: r}
}

// wrapHandler converts an abstract HandlerFunc to a gin.HandlerFunc
func (r *ginRouter) wrapHandler(handler httpmodule.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap gin.Context with our abstract Context
		ctx := newGinContext(c)

		// Call the handler
		if err := handler(ctx); err != nil {
			// If handler returns an error, abort with 500
			// Handlers can set their own status before returning error if needed
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
	}
}

// GET registers a GET route
func (r *ginRouter) GET(path string, handler httpmodule.HandlerFunc) {
	r.router.GET(path, r.wrapHandler(handler))
}

// POST registers a POST route
func (r *ginRouter) POST(path string, handler httpmodule.HandlerFunc) {
	r.router.POST(path, r.wrapHandler(handler))
}

// PUT registers a PUT route
func (r *ginRouter) PUT(path string, handler httpmodule.HandlerFunc) {
	r.router.PUT(path, r.wrapHandler(handler))
}

// DELETE registers a DELETE route
func (r *ginRouter) DELETE(path string, handler httpmodule.HandlerFunc) {
	r.router.DELETE(path, r.wrapHandler(handler))
}

// PATCH registers a PATCH route
func (r *ginRouter) PATCH(path string, handler httpmodule.HandlerFunc) {
	r.router.PATCH(path, r.wrapHandler(handler))
}

// Group creates a route group with the given prefix
func (r *ginRouter) Group(prefix string) httpmodule.Router {
	group := r.router.Group(prefix)
	return newGinRouter(group)
}
