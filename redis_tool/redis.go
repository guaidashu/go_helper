/**
  create by yy on 2019-08-29
*/

package redis_tool

import (
	"errors"
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

// init
func (r *RedisConn) Init(config *RedisConfig) error {
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
	return r.Connection()
}

// connect redis
func (r *RedisConn) Connection() error {
	conn, err := redis.Dial(r.Config.NetWork, r.Config.Address)
	if err != nil {
		return errors.New(fmt.Sprintf("redis connect error, %v", err))
	}
	if r.Config.Password != "" {
		_, err := conn.Do("auth", r.Config.Password)
		if err != nil {
			return errors.New(fmt.Sprintf("auth is failed, check your config or judge your password whether correct! error: %v", err))
		}
	}
	r.conn = conn
	return nil
}

// Set a value for a customize key
// return an interface and an error
// Origin command: GET KEY_NAME
func (r *RedisConn) Set(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("SET", *(argsP)...)
}

// Get a value by a given key
// return an interface and an error
// Origin command: GET KEY_NAME
func (r *RedisConn) Get(key string) (interface{}, error) {
	return String(r.conn.Do("GET", key))
}

// Check key whether exists
// Origin command: EXISTS KEY_NAME
func (r *RedisConn) Exists(key string) (interface{}, error) {
	return r.conn.Do("EXISTS", key)
}

// Del key
// Origin command: DEL KEY_NAME
func (r *RedisConn) Del(key string) (interface{}, error) {
	return r.conn.Do("DEL", key)
}

// Only set a value for the given key when the key is not exists
// Origin command: SETNX KEY_NAME VALUE
func (r *RedisConn) SetNX(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("SETNX", *(argsP)...)
}

// Close connection
func (r *RedisConn) Close() error {
	return r.conn.Close()
}

// Expire
// Set expire time
func (r *RedisConn) Expire(key string, timeout int) (interface{}, error) {
	return r.conn.Do("EXPIRE", key, timeout)
}

// LPUSH
// Insert one or more values into the list header
func (r *RedisConn) LPush(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LPUSH", *(argsP)...)
}

// LPUSHX
// Insert one value into the the list header
func (r *RedisConn) LPUSHX(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LPUSHX", *(argsP)...)
}

// LRANGE
//
// LRange(key, start, end)
// Start or end is beginning with 0, and -1 is meaning that the end of the list.
//
// Get the elements in the specified range of the list.
func (r *RedisConn) LRange(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LRANGE", *(argsP)...)
}

// LPOP
// Move out and get the first element of the list
func (r *RedisConn) LPop(key string) (interface{}, error) {
	return r.conn.Do("LPOP", key)
}

// LLEN
// Get list length
func (r *RedisConn) LLen(key string) (interface{}, error) {
	return r.conn.Do("LLEN", key)
}

// LINDEX
// Get the element of the list header by index.
func (r *RedisConn) LIndex(key string) (interface{}, error) {
	return r.conn.Do("LINDEX", key)
}

// LRem
// LRrm(key, count, value)
//
// count > 0: Search from the beginning of the header to the end of the table,
// removing the equivalent of value, the number is count.
//
// count < 0: Search from the beginning of the table to the header,
// remove the same element as value, the number is the absolute value of count.
//
// count = 0: Remove all values in the table that are equal to value.
func (r *RedisConn) LRem(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LREM", *(argsP)...)
}

// LSET
// LSet(key, index, value)
// Set the value of element by index.
func (r *RedisConn) LSet(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LSET", *(argsP)...)
}

// RPop
// Remove the element which in the end of the list.
// Return the value of element which is removed.
// If the index is not exists, return nil.
func (r *RedisConn) RPop(key string) (interface{}, error) {
	return r.conn.Do("RPOP", key)
}

// RPush
// RPush(key, "foo")
// Insert one or more values into the end of the list(far right).
func (r *RedisConn) RPush(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("RPUSH", *(argsP)...)
}

// RPushx
// RPushx(key, "foo")
// Insert one value into the end of the list(far right).
func (r *RedisConn) RPushx(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("RPUSHX", *(argsP)...)
}

// LINSERT
// LInsert(key, "BEFORE", element, value)
// Insert value into before the element of the list.
func (r *RedisConn) LInsert(key string, args ...interface{}) (interface{}, error) {
	argsP := go_helper.Prepend(&args, key)
	return r.conn.Do("LINSERT", *(argsP)...)
}

// convert interface to string.
func (r *RedisConn) String(value interface{}) string {
	s, _ := redis.String(value, nil)
	return s
}
