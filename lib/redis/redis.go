package redis

import (
	"strconv"
	"time"

	"net/url"

	"github.com/garyburd/redigo/redis"
	"theaterzero.com/backend/lib/config"
	"theaterzero.com/backend/lib/log"
)

var (
	redisPool   *redis.Pool
	connection  = "tcp"
	idleTimeout = time.Duration(240)
	loopSize    = 1000
)

//GetConn :nodoc:
func GetConn() redis.Conn {
	conn := redisPool.Get()
	return conn
}

//SScan :nodoc:
func SScan(key, value string) (result []string, err error) {
	conn := GetConn()

	cursor := 0
	for {
		var (
			values []interface{}
			found  []string
		)

		values, err = redis.Values(conn.Do("SSCAN", key, cursor, "MATCH", "*"+value+"*", "COUNT", loopSize))
		if err != nil {
			log.Error(err)
			return
		}

		found, err = redis.Strings(values[1], err)
		if err != nil {
			log.Error(err)
			return
		}

		if len(found) > 0 {
			result = append(result, found...)
		}

		cursor, err = redis.Int(values[0], err)
		if err != nil {
			log.Error(err)
			return
		}
		if cursor == 0 {
			break
		}
	}

	return
}

// SAdd :nodoc:
func SAdd(key, value string) (err error) {
	conn := GetConn()

	_, err = conn.Do("SADD", key, value)
	if err != nil {
		return err
	}

	return
}

// Del :nodoc:
func Del(key string) (err error) {
	conn := GetConn()

	_, err = conn.Do("DEL", key)
	if err != nil {
		return err
	}

	return
}

func init() {
	r, err := url.Parse(config.RedisURL())
	if err != nil {
		log.Fatal(err)
	}

	var password string
	if r.User != nil {
		if pwd, ok := r.User.Password(); ok {
			password = pwd
		} else {
			password = r.User.Username()
		}
	}

	var database string
	if r.Path != "" {
		db, err := strconv.ParseInt(r.Path, 10, 64)
		if err == nil {
			database = strconv.Itoa(int(db))
		}
	}

	redisPool = &redis.Pool{
		MaxIdle:     config.RedisPoolCount(),
		IdleTimeout: idleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(connection, r.Host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if database != "" {
				if _, err := c.Do("SELECT", database); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}
