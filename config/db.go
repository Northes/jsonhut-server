package config

var DB dbStruct

type dbStruct struct {
	Addr       string `ini:"database_addr"`
	Port       string `ini:"database_port"`
	DBName     string `ini:"database_name"`
	DBUserName string `ini:"database_username"`
	DBPassword string `ini:"database_password"`
}
