package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

type RedisConnector struct {
	Db redis.Conn
}

type Result struct {
	Valid bool
	Foo   string
	FF    int
}

func (r *RedisConnector) Connect() {
	c, err := redis.Dial("tcp", ":6379")
	r.Db = c
	if err != nil {
		panic(err)
	}
}

func (r *RedisConnector) Close() {
	r.Db.Close()
}

func (r *RedisConnector) Get(key string) Result {
	inter, err := r.Db.Do("get", "foo")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get from redis:", inter)
	fmt.Println("Get type:", reflect.TypeOf(inter))

	res := Result{Valid: false}
	if inter == nil {
		return res
	}
	bytes, _ := inter.([]byte)

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		panic(err)
	}

	res.Valid = true
	return res
}

func (r *RedisConnector) Rpush(key string, value string) int64 {
	result, err := r.Db.Do("rpush", key, value)
	if err != nil {
		panic(err)
	}
	return result.(int64)
}
