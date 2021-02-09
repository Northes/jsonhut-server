package config

var Redis redisStruct

//redis配置结构
type redisStruct struct {
	Addr string `ini:"redis_addr"`
}