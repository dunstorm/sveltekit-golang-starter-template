package config

// Config holds all configuration for the application
type Config struct {
	HTTP          HTTPConfig
	IsDevelopment bool
}

// HTTPConfig holds HTTP-specific configuration
type HTTPConfig struct {
	Port         string
	APIKeySecret string
}

// New creates a new configuration with default values
func New() *Config {
	return &Config{
		HTTP: HTTPConfig{
			Port: "8080",
		},
		IsDevelopment: true,
	}
}

// IsDev returns whether the application is running in development mode
func (c *Config) IsDev() bool {
	return c.IsDevelopment
}
