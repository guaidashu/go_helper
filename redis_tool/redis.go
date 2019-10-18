/**
  create by yy on 2019-08-29
*/

package redis_tool

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/guaidashu/go_helper"
)

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
	NetWork  string `json:"net_work"`
	Address  string `json:"address"`
}

type RedisConn struct {
	conn   redis.Conn
	Config *RedisConfig
}

/**
init
*/
func (r *RedisConn) Init(config *RedisConfig) {
	if config == nil {
		r.Config = &RedisConfig{
			Host:     "127.0.0.1",
			Password: "",
			Port:     "6379",
			NetWork:  "tcp",
		}
	} else {
		if config.NetWork == "" {
			config.NetWork = "tcp"
		}
		if config.Host == "" {
			config.Host = "127.0.0.1"
		}
		if config.Port == "" {
			config.Port = "6379"
		}
		if config.Address == "" {
			config.Address = fmt.Sprintf("%v:%v", config.Host, config.Port)
		}
		r.Config = config
	}
	r.Connection()
}

/**
connect redis
*/
func (r *RedisConn) Connection() {
	conn, err := redis.Dial(r.Config.NetWork, r.Config.Address)
	if err != nil {
		panic(fmt.Sprintf("redis connect error, %v", err))
	}
	if r.Config.Password != "" {
		_, err := conn.Do("auth", r.Config.Password)
		if err != nil {
			panic(fmt.Sprintf("auth is failed, check your config or judge your password whether correct! error: %v", err))
		}
	}
	r.conn = conn
}

/**
Set a value for a customize key
return an interface and an error
Origin command: GET KEY_NAME
*/
func (r *RedisConn) Set(key string, args ...interface{}) (interface{}, error) {
	args = *(go_helper.Prepend(&args, key))
	return r.conn.Do("SET", args...)
}

/**
Get a value by a given key
return an interface and an error
Origin command: GET KEY_NAME
*/
func (r *RedisConn) Get(key string) (interface{}, error) {
	return String(r.conn.Do("GET", key))
}

/**
Check key whether exists
Origin command: EXISTS KEY_NAME
*/
func (r *RedisConn) Exists(key string) (interface{}, error) {
	return r.conn.Do("EXISTS", key)
}

/**
Del key
Origin command: DEL KEY_NAME
*/
func (r *RedisConn) Del(key string) (interface{}, error) {
	return r.conn.Do("DEL", key)
}

/**
Only set a value for the given key when the key is not exists
Origin command: SETNX KEY_NAME VALUE
*/
func (r *RedisConn) SetNX(key string, args ...interface{}) (interface{}, error) {
	args = *(go_helper.Prepend(&args, key))
	return r.conn.Do("SETNX", args...)
}

/**
Close connection
*/
func (r *RedisConn) Close() error {
	return r.conn.Close()
}

/**
Set expire time
*/
func (r *RedisConn) Expire(key string, timeout int) (interface{}, error) {
	return r.conn.Do("EXPIRE", key, timeout)
}

/**
lpush
*/
func (r *RedisConn) LPush(key string, args ...interface{}) (interface{}, error) {
	args = *(go_helper.Prepend(&args, key))
	return r.conn.Do("lpush", args...)
}

/**
lrange
*/
func (r *RedisConn) LRange(key string, args ...interface{}) (interface{}, error) {
	args = *(go_helper.Prepend(&args, key))
	return r.conn.Do("lrange", args...)
}

func (r *RedisConn) String(value interface{}) string {
	s, _ := redis.String(value, nil)
	return s
}
