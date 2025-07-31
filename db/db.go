package db

import (
	"context"
	"net/url"
	"time"

	"github.com/KDany/gotils/log"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	StartConnect  string
	RetryConnect  string
	FailedConnect string
	Connected     string
	Disconnect    string
}

type loggerImpl struct{}

var conf Config

func (loggerImpl) SetConfig(c Config) {
	conf = c
}

var Logger loggerImpl

var baseConfig = Config{
	StartConnect:  "/gotils/ - Connecting to database at ",
	RetryConnect:  "/gotils/ - Failed to connect to database. Retrying in 5 seconds...",
	FailedConnect: "/gotils/ - Failed to connect to database: ",
	Connected:     "/gotils/ - Database connection established",
	Disconnect:    "/gotils/ - Disconnecting from database",
}

func getConfig() Config {
	if conf == (Config{}) {
		return baseConfig
	}
	return conf
}

var db *pgx.Conn

func Get() *pgx.Conn {
	return db
}

func Connect(u string, retries int) error {
	var err error
	var logConfig = getConfig()

	log.Info(logConfig.StartConnect, maskDBUrl(u))

	for ; retries > 0; retries-- {
		db, err = pgx.Connect(context.Background(), u)
		if err == nil {
			break
		}
		log.Warn(logConfig.RetryConnect)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Error(logConfig.FailedConnect, err)
		return err
	}
	log.Info(logConfig.Connected)
	return nil
}

func Disconnect() error {
	if db != nil {
		var logConfig = getConfig()
		log.Info(logConfig.Disconnect)
		return db.Close(context.Background())
	}
	return nil
}

func Ping() error {
	if db != nil {
		return db.Ping(context.Background())
	}
	return nil
}

func IsConnected() bool {
	if db != nil {
		err := db.Ping(context.Background())
		return err == nil
	}
	return false
}

func maskDBUrl(rawurl string) string {
	parsed, err := url.Parse(rawurl)
	if err != nil {
		return rawurl
	}
	if parsed.User != nil {
		parsed.User = url.UserPassword("******", "******")
		masked := parsed.Scheme + "://"
		masked += "******:******@"
		if parsed.Host != "" {
			masked += parsed.Host
		}
		if parsed.Path != "" {
			masked += parsed.Path
		}
		if parsed.RawQuery != "" {
			masked += "?" + parsed.RawQuery
		}
		if parsed.Fragment != "" {
			masked += "#" + parsed.Fragment
		}
		return masked
	}
	return parsed.String()
}
