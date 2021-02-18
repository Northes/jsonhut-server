package config

import (
	"gopkg.in/ini.v1"
)

var App appStruct

type appStruct struct {
	AppName               string `ini:"app_name"`
	Port                  string `ini:"app_port"`
	BaseUrl               string `ini:"base_url"`
	PostIPCurrentLimit    int    `ini:"post_ip_current_limit"`
	GetIPCurrentLimit     int    `ini:"get_ip_current_limit"`
	DetailsIPCurrentLimit int    `ini:"details_ip_current_limit"`
}

type configStruct struct {
	App   appStruct   `ini:"app"`
	Redis redisStruct `ini:"redis"`
	DB    dbStruct    `ini:"database"`
}

func LoadConfig() {
	//cfgDefault, err := ini.Load(".env")
	//if err != nil {
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exit(1)
	//}
	//
	//var cfg *ini.File
	//if cfgDefault.Section("").Key("app_mode").String() == "development" {
	//	cfg, err = ini.Load(".env.dev")
	//} else {
	//	cfg, err = ini.Load(".env.prod")
	//}
	//if err != nil {
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exit(1)
	//}
	cfg, err := ini.LooseLoad(".env", ".env.local")
	if err != nil {
		panic(err)
	}
	config := new(configStruct)

	err = cfg.MapTo(config)
	if err != nil {
		panic(err)
	}
	App = config.App
	Redis = config.Redis
	//Log = config.Log
	DB = config.DB
}
