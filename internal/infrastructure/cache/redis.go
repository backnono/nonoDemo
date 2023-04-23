package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"nonoDemo/pkg/framework"
	redisCfg "nonoDemo/pkg/framework/redis"
	"nonoDemo/pkg/utils"
)

type Connection struct {
	client *redis.Client
}

func NewConnection(cfg framework.Configuration) *Connection {
	config := utils.LoadConfig(cfg, redisCfg.Config{}).(redisCfg.Config)
	rd := &Connection{}
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	options := &redis.Options{
		Network: "tcp",
		Addr:    addr,
		// Password: config.Password,
		DB: config.DB,
		// PoolSize:     config.PoolSize,
		// MinIdleConns: config.IdleConns,
	}
	if config.Password != "" {
		options.Password = config.Password
	}
	rd.client = redis.NewClient(options)
	_, err := rd.client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return rd
}

func (rd *Connection) String(key string) (string, error) {
	return rd.client.Get(key).Result()
}

func (rd *Connection) Get(key string, value interface{}) error {
	// str, err := rd.String(key)
	// if err != nil {
	// 	return err
	// }
	// err = value.(encoding.BinaryUnmarshaler).UnmarshalBinary([]byte(str))
	// return err
	return rd.client.Get(key).Scan(value)
}

func (rd *Connection) Int(key string) (int, error) {
	return rd.client.Get(key).Int()
}
func (rd *Connection) Int64(key string) (int64, error) {
	return rd.client.Get(key).Int64()
}

func (rd *Connection) HGetAll(key string) (map[string]string, error) {
	return rd.client.HGetAll(key).Result()
}

func (rd *Connection) HString(key, field string) (
	string, error) {
	return rd.client.HGet(key, field).Result()
}

func (rd *Connection) HInt(key, field string) (int, error) {
	return rd.client.HGet(key, field).Int()
}

func (rd *Connection) Set(key string, value interface{},
	expired int) error {
	return rd.client.Set(key, value,
		time.Duration(expired)*time.Second).Err()
}

func (rd *Connection) Del(key string) error {
	return rd.client.Del(key).Err()
}

func (rd *Connection) HSet(key, field string, value interface{}) error {
	err := rd.client.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	rd.client.Expire(key, time.Hour*30*24)
	return rd.client.HSet(key, field, value).Err()
}
func (rd *Connection) HDelete(key, field string) error {
	return rd.client.HDel(key, field).Err()
}

func (rd *Connection) SMembers(key string) ([]string, error) {
	return rd.client.SMembers(key).Result()
}

func (rd *Connection) Incr(key string, factor int64) (int64, error) {
	return rd.client.IncrBy(key, factor).Result()
}

func (rd *Connection) Expire(key string, expired int) (bool, error) {
	return rd.client.Expire(key, time.Duration(expired)*time.Second).Result()
}

func (rd *Connection) Exist(key string) (bool, error) {
	if cnt, err := rd.client.Exists(key).Result(); err != nil {
		return false, err
	} else {
		return cnt > 0, nil
	}
}

func (rd *Connection) ExecuteScript(script string, keys []string, args ...string) (interface{}, error) {
	return rd.client.Eval(script, keys, args).Result()
}

func (rd *Connection) Close() {
	rd.client.Close()
}

func (rd *Connection) Connection() *redis.Client {
	return rd.client
}

func (rd *Connection) Keys(prefix string) ([]string, error) {
	return rd.client.Keys(prefix).Result()
}

func (rd *Connection) DelMultiple(key []string) error {
	return rd.client.Del(key...).Err()
}

// ScanLimit
// @description: 通过游标扫描，替换全量key读取的慢查询情况
// @param {string} match
// @param {int64} window
// @param {int} limit
// @return {*}
func (rd *Connection) ScanLimit(match string, window int64, limit int) ([]string, error) {
	// 高并发场景下，右边读取可能存在key重复的情况！此处逻辑没有去重！
	var (
		cursor uint64
		keys   []string
		n      int
		err    error
	)
	for {
		var eachKeys []string
		eachKeys, cursor, err = rd.client.Scan(cursor, match, window).Result()
		if err != nil {
			return keys, err
		} else {
			keys = append(keys, eachKeys...)
			n += len(eachKeys)
		}
		if limit != 0 && n >= limit {
			return keys[:limit], nil
		}
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}
