package httpgin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/things-kit/core/log"
)

// GinServer implements the http.Server interface using Gin.
type GinServer struct {
	engine *gin.Engine
	server *http.Server
	config *Config
	logger log.Logger
}

// NewGinServer creates a new Gin server instance.
func NewGinServer(config *Config, logger log.Logger) *GinServer {
	// Set Gin mode
	gin.SetMode(config.Mode)

	// Create Gin engine
	engine := gin.New()

	// Add default middleware
	engine.Use(gin.Recovery())

	return &GinServer{
		engine: engine,
		config: config,
		logger: logger,
	}
}

// Start implements http.Server.Start
func (s *GinServer) Start(ctx context.Context) error {
	addr := s.Addr()
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	s.logger.Info("Starting Gin HTTP server", log.Field{Key: "address", Value: addr})

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("Gin HTTP server error", err, log.Field{Key: "address", Value: addr})
		}
	}()

	return nil
}

// Stop implements http.Server.Stop
func (s *GinServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping Gin HTTP server", log.Field{Key: "address", Value: s.Addr()})

	if s.server == nil {
		return nil
	}

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(shutdownCtx)
}

// Addr implements http.Server.Addr
func (s *GinServer) Addr() string {
	if s.config.Config.Host != "" {
		return fmt.Sprintf("%s:%d", s.config.Config.Host, s.config.Config.Port)
	}
	return fmt.Sprintf(":%d", s.config.Config.Port)
}

// Engine returns the underlying Gin engine.
// This is useful for registering middleware or customizing the engine.
func (s *GinServer) Engine() *gin.Engine {
	return s.engine
}
