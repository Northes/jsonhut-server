package dao

import (
	"fmt"
	"jsonhut-server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := config.DB.DBUserName + ":" + config.DB.DBPassword + "@tcp(" + config.DB.Addr + ")/" + config.DB.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "jsonhut:y8Wx4ZkMXjnAnMfz@tcp(127.0.0.1:3306)/jsonhut?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//err2 := DB.AutoMigrate(&models.Json{})
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	return nil
}
