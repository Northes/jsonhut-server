package models

import (
	"fmt"
	"gorm.io/gorm"
	"jsonhutapi/dao"
	"jsonhutapi/logic"
	"time"
)

type Json struct {
	gorm.Model
	JsonId         string    `gorm:"type:varchar(255);uniqueIndex:idx_jsons_json_id"` // JsonID
	JsonBody       string    // Json主体
	ExpirationTime time.Time `gorm:"type:time"` // 过期时间
	CallCount      uint      `gorm:"type:uint"` // 调用次数
	Status         uint      `gorm:"type:uint"` // 状态：0 正常 1 禁用 2 审核中 3 审核拒绝
	FromIP         string    // 来源IP
	Comment        string    // 备注
}

func CreateAJson(jsonBody string, expirationTime int) (uint, error) {
	//db := Connect2DataBase()
	// 设置json过期时间
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
		ExpirationTime: time.Unix(nowTime, 0),
		CallCount:      0,
		Status:         0,
		FromIP:         "",
		Comment:        "",
	}
	// 入库
	result := dao.DB.Create(&json)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, result.Error
	}

	return json.ID, nil
	//fmt.Println(result)
}

func QueryJsonBodyByJsonID(jsonID string) (Json, error) {
	//db := Connect2DataBase()

	//var result map[string]interface{}
	//db.Model(&Json{}).First(&result, "json_id = ?", jsonID)
	//result := db.Where(&Json{JsonId: jsonID}).Select("json_body").First(&Json{})

	//result := map[string]interface{}{}
	var json Json
	result := dao.DB.Model(&Json{}).First(&json, "json_id = ?", jsonID)
	//fmt.Println(result["json_body"])
	if result.Error != nil {
		fmt.Println(result.Error)
		return json, result.Error
	}

	//fmt.Println(json.JsonBody)

	return json, nil
}

func UpdateJsonID(id uint) string {
	//db := Connect2DataBase()
	jsonId := logic.Encode(uint64(id))
	dao.DB.Model(&Json{}).Where("id = ?", id).Update("json_id", jsonId)
	return jsonId
}

func UpdateJsonCallCount(jsonID string) {
	resultData, err := QueryJsonBodyByJsonID(jsonID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dao.DB.Model(&Json{}).Where("json_id = ?", jsonID).Update("call_count", resultData.CallCount+1)
}
