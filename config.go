package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	httpmodule "github.com/things-kit/things-kit-http"
)

// Config holds the Gin-specific HTTP server configuration.
// It embeds the common http.Config and adds Gin-specific options.
type Config struct {
	httpmodule.Config `mapstructure:",squash"`
	Mode              string `mapstructure:"mode"` // debug, release, test
}

// NewConfig creates a new Gin HTTP configuration from Viper.
func NewConfig(v *viper.Viper) *Config {
	cfg := &Config{
		Config: httpmodule.Config{
			Port: 8080,
			Host: "",
		},
		Mode: gin.ReleaseMode,
	}

	// Load configuration from viper
	if v != nil {
		_ = v.UnmarshalKey("http", cfg)
	}

	return cfg
}
