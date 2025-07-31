package db

import (
	"context"
	"net/url"
	"time"

	"github.com/KDany/gotils/log"
	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func Get() *pgx.Conn {
	return db
}

func Connect(u string, retries int) error {
	var err error

	log.Info("/gotils/ - Connecting to database at ", maskDBUrl(u))

	for ; retries > 0; retries-- {
		db, err = pgx.Connect(context.Background(), u)
		if err == nil {
			break
		}
		log.Warn("/gotils/ - Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Error("/gotils/ - Failed to connect to database: ", err)
		return err
	}
	return nil
}

func Disconnect() error {
	if db != nil {
		log.Info("/gotils/ - Disconnecting from database")
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
