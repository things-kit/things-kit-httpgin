// Package httpgin provides a Gin-based implementation of the Things-Kit HTTP server interface.
// This is the default HTTP implementation for Things-Kit, but users can provide their own
// implementations using different frameworks (Chi, Echo, stdlib, etc.).
package httpgin

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/things-kit/core/log"
	httpmodule "github.com/things-kit/things-kit-http"
	"go.uber.org/fx"
)

// Module provides the Gin-based HTTP server module to the application.
// This module implements the http.Server interface using the Gin framework.
var Module = fx.Module("httpgin",
	fx.Provide(
		NewConfig,
		NewGinServer,
		fx.Annotate(
			func(s *GinServer) httpmodule.Server { return s },
			fx.As(new(httpmodule.Server)),
		),
	),
	fx.Invoke(RunHttpServer),
)

// GinHandler is a Gin-specific implementation of http.Handler.
// Handlers that implement this interface can be registered with AsGinHandler.
type GinHandler interface {
	RegisterRoutes(engine *gin.Engine)
}

// HttpServerParams contains all dependencies needed to run the HTTP server.
type HttpServerParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    log.Logger
	Config    *Config
	Handlers  []GinHandler `group:"http.handlers"`
}

// RunHttpServer starts the HTTP server with registered handlers.
// This is invoked by Fx during application startup.
func RunHttpServer(p HttpServerParams, server *GinServer) {
	// Register all provided handlers
	for _, handler := range p.Handlers {
		handler.RegisterRoutes(server.engine)
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return server.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}

// AsGinHandler is a generic helper to provide a Gin HTTP handler to the Fx graph.
// The constructor should return a type that implements the GinHandler interface.
//
// Example:
//
//	type MyHandler struct {
//	    logger log.Logger
//	}
//
//	func NewMyHandler(logger log.Logger) *MyHandler {
//	    return &MyHandler{logger: logger}
//	}
//
//	func (h *MyHandler) RegisterRoutes(engine *gin.Engine) {
//	    engine.GET("/hello", func(c *gin.Context) {
//	        c.JSON(200, gin.H{"message": "Hello World"})
//	    })
//	}
//
//	// In main.go:
//	httpgin.AsGinHandler(NewMyHandler)
func AsGinHandler(constructor any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			constructor,
			fx.As(new(GinHandler)),
			fx.ResultTags(`group:"http.handlers"`),
		),
	)
}
