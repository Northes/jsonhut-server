package dao

import (
	"fmt"
	"jsonhut-server/config"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool //创建redis连接池
)

func InitRedis() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", config.Redis.Addr)
		},
	}
}

func RedisSetData(name string, data string, exTime int) {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("Set", name, data, "EX", exTime)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RedisGetData(name string) (string, error) {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	r, err := redis.String(c.Do("GET", name))
	if err != nil {
		fmt.Println("Redis Get "+name+" faild :", err)
		return "", err
	}
	return r, nil
}

func RedisSetExpirationTime(name string, exTime int) {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池
	// 如果过期时间设为-1，则重置为默认5分钟
	if exTime == -1 {
		exTime = 5 * 60
	}
	_, err := c.Do("expire", name, exTime)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RedisGetTTL(name string) int {
	c := pool.Get()
	defer c.Close()

	ttl, err := redis.Int(c.Do("TTL", name))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(ttl)
	return ttl
}
