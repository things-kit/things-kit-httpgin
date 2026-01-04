package httpgin

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	httpmodule "github.com/things-kit/things-kit-http"
)

// ginContext wraps gin.Context to implement the http.Context interface.
// This allows handlers to be framework-agnostic while still using Gin under the hood.
type ginContext struct {
	ctx *gin.Context
}

// newGinContext creates a new wrapper around gin.Context.
func newGinContext(c *gin.Context) httpmodule.Context {
	return &ginContext{ctx: c}
}

// Request returns the underlying *http.Request
func (c *ginContext) Request() *http.Request {
	return c.ctx.Request
}

// Context returns the request's context for cancellation and deadlines
func (c *ginContext) Context() context.Context {
	return c.ctx.Request.Context()
}

// Param retrieves a URL path parameter by name
func (c *ginContext) Param(name string) string {
	return c.ctx.Param(name)
}

// Query retrieves a URL query parameter by name
func (c *ginContext) Query(name string) string {
	return c.ctx.Query(name)
}

// QueryDefault retrieves a URL query parameter with a default value
func (c *ginContext) QueryDefault(name, defaultValue string) string {
	return c.ctx.DefaultQuery(name, defaultValue)
}

// GetHeader retrieves a request header by name
func (c *ginContext) GetHeader(name string) string {
	return c.ctx.GetHeader(name)
}

// SetHeader sets a response header
func (c *ginContext) SetHeader(name, value string) {
	c.ctx.Header(name, value)
}

// BindJSON binds the request body as JSON to the provided struct
func (c *ginContext) BindJSON(obj interface{}) error {
	return c.ctx.BindJSON(obj)
}

// Bind binds the request body to the provided struct (supports multiple formats)
func (c *ginContext) Bind(obj interface{}) error {
	return c.ctx.Bind(obj)
}

// JSON sends a JSON response with the given status code
func (c *ginContext) JSON(code int, obj interface{}) error {
	c.ctx.JSON(code, obj)
	return nil
}

// String sends a string response with the given status code
func (c *ginContext) String(code int, s string) error {
	c.ctx.String(code, s)
	return nil
}

// Status sets the HTTP response status code
func (c *ginContext) Status(code int) {
	c.ctx.Status(code)
}

// Writer returns the response writer
func (c *ginContext) Writer() io.Writer {
	return c.ctx.Writer
}
