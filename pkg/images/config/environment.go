package config

import (
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
)

type LogLevel string

type EnvironmentConfig struct {
	// TrustProxy sets whether HTTP headers from proxy are to be trusted
	TrustProxy config.TrustProxy `envconfig:"TRUST_PROXY" default:"false"`
	// LogLevel sets the global log level
	LogLevel LogLevel `envconfig:"LOG_LEVEL" default:"warn"`
	// SentryDSN sets the sentry DSN.
	SentryDSN config.SentryDSN `envconfig:"SENTRY_DSN"`
	// ConfigSource configures the source of app configurations
	ConfigSource *configsource.Config `envconfig:"CONFIG_SOURCE"`
	// BuiltinResourceDirectory sets the directory for built-in resource files
	BuiltinResourceDirectory string `envconfig:"BUILTIN_RESOURCE_DIRECTORY" default:"resources/authgear"`
	// CustomResourceDirectory sets the directory for customized resource files
	CustomResourceDirectory string `envconfig:"CUSTOM_RESOURCE_DIRECTORY"`
	// GlobalDatabase configures the global database for configuresource
	GlobalDatabase *config.GlobalDatabaseCredentialsEnvironmentConfig `envconfig:"DATABASE"`
	// DatabaseConfig configures the database connection config
	DatabaseConfig *config.DatabaseEnvironmentConfig `envconfig:"DATABASE_CONFIG"`
	// CORSAllowOrigins configures a comma-separated list of allowed origins for CORSMiddleware
	CORSAllowedOrigins config.CORSAllowedOrigins `envconfig:"CORS_ALLOWED_ORIGINS"`
}
