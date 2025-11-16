package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// This package centralizes how every service reads configuration. The pattern is:
//   * Define strongly typed sections (HTTP, Database, Cache, etc.).
//   * Populate them via `Load`, which checks for <SERVICE>_ overrides and then falls
//     back to global env vars or sane defaults (see defaultPorts below).
//   * Pass the resulting ServiceConfig down through bootstrap so the rest of the
//     layers can focus on business logic rather than env parsing.

var defaultPorts = map[string]int{
	"gateway":      9000,
	"auth":         9101,
	"blog":         9102,
	"interaction":  9103,
	"notification": 9104,
	"analytics":    9105,
	"bookmark":     9106,
	"search":       9107,
	"worker":       9201,
}

// ServiceConfig bundles the configuration shared by each service.
type ServiceConfig struct {
	Name          string
	Environment   string
	HTTP          HTTPConfig
	Database      DatabaseConfig
	Cache         CacheConfig
	Messaging     MessagingConfig
	Observability ObservabilityConfig
}

// HTTPConfig defines listener level settings.
type HTTPConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (h HTTPConfig) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

// DatabaseConfig captures relational database connectivity.
type DatabaseConfig struct {
	URL           string
	MaxConns      int
	MigrationsDir string
}

// CacheConfig holds redis / in-memory cache settings.
type CacheConfig struct {
	URL      string
	Username string
	Password string
}

// MessagingConfig configures the async event bus.
type MessagingConfig struct {
	BrokerURL   string
	Exchange    string
	ConsumerTag string
}

// ObservabilityConfig surface metrics + tracing knobs.
type ObservabilityConfig struct {
	PrometheusPort int
	TraceEndpoint  string
}

// Load assembles a ServiceConfig using environment overrides.
//
// Order of precedence for each setting:
//   1. <SERVICE>_ scoped env var (e.g. AUTH_HTTP_PORT).
//   2. Global env var (e.g. HTTP_PORT) if available.
//   3. Hard-coded defaults (ports live in defaultPorts; others inline below).
//
// The map lookups + helpers at the bottom keep all parsing (ints/durations) in one
// place which simplifies validation/testing later.
func Load(service string) ServiceConfig {
	port := defaultPorts[service]
	if port == 0 {
		port = 9100
	}

	envPrefix := fmt.Sprintf("%s_", strings.ToUpper(service))

	return ServiceConfig{
		Name:        service,
		Environment: getenv("APP_ENV", "development"),
		HTTP: HTTPConfig{
			Host:         getenv(envPrefix+"HTTP_HOST", getenv("HTTP_HOST", "0.0.0.0")),
			Port:         getenvInt(envPrefix+"HTTP_PORT", getenvInt("HTTP_PORT", port)),
			ReadTimeout:  getenvDuration(envPrefix+"HTTP_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getenvDuration(envPrefix+"HTTP_WRITE_TIMEOUT", 15*time.Second),
		},
		Database: DatabaseConfig{
			URL:           getenv(envPrefix+"DATABASE_URL", os.Getenv("DATABASE_URL")),
			MaxConns:      getenvInt(envPrefix+"DB_MAX_CONNS", 5),
			MigrationsDir: getenv(envPrefix+"MIGRATIONS_DIR", "migrations"),
		},
		Cache: CacheConfig{
			URL:      getenv(envPrefix+"CACHE_URL", os.Getenv("CACHE_URL")),
			Username: getenv(envPrefix+"CACHE_USER", ""),
			Password: getenv(envPrefix+"CACHE_PASSWORD", ""),
		},
		Messaging: MessagingConfig{
			BrokerURL:   getenv(envPrefix+"BROKER_URL", os.Getenv("BROKER_URL")),
			Exchange:    getenv(envPrefix+"BROKER_EXCHANGE", "portfolio"),
			ConsumerTag: getenv(envPrefix+"CONSUMER_TAG", service+"-consumer"),
		},
		Observability: ObservabilityConfig{
			PrometheusPort: getenvInt(envPrefix+"PROM_PORT", port+100),
			TraceEndpoint:  getenv(envPrefix+"TRACE_ENDPOINT", os.Getenv("TRACE_ENDPOINT")),
		},
	}
}

func getenv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

func getenvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			return parsed
		}
	}
	return fallback
}

func getenvDuration(key string, fallback time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if parsed, err := time.ParseDuration(val); err == nil {
			return parsed
		}
	}
	return fallback
}
