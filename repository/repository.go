package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Json struct {
	gorm.Model
	JsonId         string // JsonID
	JsonBody       string // Json主体
	ExpirationTime string `gorm:"type:time"` // 过期时间
	CallCount      uint   `gorm:"type:uint"` // 调用次数
	Status         uint   `gorm:"type:uint"` // 状态：0 正常 1 禁用 2 审核中 3 审核拒绝
	FromIP         string // 来源IP
	Comment        string // 备注
}

func Connect2DataBase() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jsonhut?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "jsonhut:y8Wx4ZkMXjnAnMfz@tcp(127.0.0.1:3306)/jsonhut?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err2 := db.AutoMigrate(&Json{})
	if err2 != nil {
		fmt.Println(err2)
	}
	return db
}

func CreateJson(jsonBody string, expirationTime int) (uint, error) {
	db := Connect2DataBase()
	nowTime := time.Now().Unix()
	switch expirationTime {
	case 0:
		nowTime = 32472115200
		break
	case 1:
		nowTime += 86400
		break
	case 3:
		nowTime += 259200
		break
	case 7:
		nowTime += 604800
		break
	default:
		nowTime += 259200
	}

	json := Json{
		Model:          gorm.Model{},
		JsonId:         "",
		JsonBody:       jsonBody,
		ExpirationTime: time.Unix(nowTime, 0).Format("2006-01-02 15:04:05"),
		CallCount:      0,
		Status:         0,
		FromIP:         "",
		Comment:        "",
	}
	result := db.Create(&json)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, result.Error
	}

	return json.ID, nil
	//fmt.Println(result)
}

func QueryJsonBodyByJsonID(jsonID string) (Json, error) {
	db := Connect2DataBase()
	//var result map[string]interface{}
	//db.Model(&Json{}).First(&result, "json_id = ?", jsonID)
	//result := db.Where(&Json{JsonId: jsonID}).Select("json_body").First(&Json{})

	//result := map[string]interface{}{}
	var json Json
	result := db.Model(&Json{}).First(&json, "json_id = ?", jsonID)
	//fmt.Println(result["json_body"])
	if result.Error != nil {
		fmt.Println(result.Error)
		return json, result.Error
	}

	//fmt.Println(json.JsonBody)

	return json, nil
}

func UpdateJsonID(id uint) string {
	db := Connect2DataBase()
	jsonId := Encode(uint64(id))
	db.Model(&Json{}).Where("id = ?", id).Update("json_id", jsonId)
	return jsonId
}

const (
	BASE    = "E8S2DZX9WYLTN6BQF7CP5IK3MJUAR4HV"
	DECIMAL = 32
	PAD     = "G"
	LEN     = 6
)

func Encode(uid uint64) string {
	id := uid
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}
