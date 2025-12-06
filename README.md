# Things-Kit HTTP Gin

**Gin-based HTTP Server for Things-Kit**

This is the default HTTP server implementation for Things-Kit, built on the [Gin framework](https://github.com/gin-gonic/gin).

## Installation

```bash
go get github.com/things-kit/things-kit-httpgin
```

## Features

- Automatic HTTP server lifecycle management
- Gin framework integration
- Configuration via Viper or environment variables
- Graceful startup and shutdown
- Health check endpoint

## Quick Start

```go
package main

import (
    "github.com/things-kit/things-kit/app"
    "github.com/things-kit/things-kit/logging"
    "github.com/things-kit/things-kit/viperconfig"
    "github.com/things-kit/things-kit-httpgin"
    httpmodule "github.com/things-kit/things-kit-http"
)

func main() {
    app.New(
        viperconfig.Module,
        logging.Module,
        httpgin.Module,
        fx.Invoke(RegisterRoutes),
    ).Run()
}

func RegisterRoutes(server httpmodule.Server) {
    engine := server.GetEngine().(*gin.Engine)
    
    engine.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, World!"})
    })
}
```

## Configuration

Via `config.yaml`:
```yaml
http:
  port: 8080
  mode: debug  # or "release"
```

Or environment variables:
```bash
export HTTP_PORT=8080
export HTTP_MODE=release
```

## Built-in Endpoints

- `GET /health` - Health check endpoint (always returns 200 OK)

## Advanced Usage

Access the Gin engine directly for middleware and advanced routing:

```go
func RegisterMiddleware(server httpmodule.Server) {
    engine := server.GetEngine().(*gin.Engine)
    
    engine.Use(gin.Recovery())
    engine.Use(gin.Logger())
    
    // Your custom middleware
    engine.Use(func(c *gin.Context) {
        // middleware logic
        c.Next()
    })
}
```

## License

MIT License - see LICENSE file for details
