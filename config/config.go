package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var App appStruct

type appStruct struct {
	AppName string `ini:"app_name"`
	Port    string `ini:"app_port"`
}

type configStruct struct {
	App   appStruct   `ini:"app"`
	Redis redisStruct `ini:"redis"`
	DB    dbStruct    `ini:"database"`
}

func LoadConfig() {
	cfgDefault, err := ini.Load(".env")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	var cfg *ini.File
	if cfgDefault.Section("").Key("app_mode").String() == "development" {
		cfg, err = ini.Load(".env.dev")
	} else {
		cfg, err = ini.Load(".env.prod")
	}
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
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
