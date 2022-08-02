package config

import (
	"log"
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

//PostgresSQL config
type PostgresSQL struct {
	Address         string        `env:"PG_ADDR" envDefault:"127.0.0.1:5432"`
	User            string        `env:"PG_USER" envDefault:"mkit"`
	Password        string        `env:"PG_PWD" envDefault:"mkit"`
	DatabaseName    string        `env:"PG_DB" envDefault:"mkit"`
	MaxOpenConns    int           `env:"PG_MAX_OPEN_CONNS" envDefault:"50"`
	MaxIdleConns    int           `env:"PG_MAX_IDLE_CONNS" envDefault:"5"`
	ConnMaxLifetime time.Duration `env:"PG_CONN_MAX_LIFETIME" envDefault:"5m"`
}

type Logger struct {
	Level    string `env:"LOGGER_LEVEL"  envDefault:"info"`
	Debug    bool   `env:"LOGGER_DEBUG"  envDefault:"1"`
	Encoding string `env:"LOGGER_ENCODING"  envDefault:"console"`
}

type Keeper struct {
	UseCache            bool `env:"USE_CACHE"  envDefault:"0"`
	WorkersProcessEvent int  `env:"WORKERS_PROCESS_EVENT"  envDefault:"3"`
}

type Metrics struct {
	Port        string
	URL         string
	ServiceName string
}

// GRPC gRPC service config
type GRPC struct {
	Port              string
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
}

// HTTP server config
type HTTP struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Config struct {
	Logger      Logger
	PostgresSQL PostgresSQL
	Keeper      Keeper
	Metrics     Metrics
	GRPC        GRPC
	HTTP        HTTP
}

var (
	config Config
	once   sync.Once
)

func ParseConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()
		if err := env.Parse(&config); err != nil {
			log.Fatal(err)
		}
	})
	return &config

}
