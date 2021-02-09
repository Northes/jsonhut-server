package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	BASE    = "E8S2DZX9WYLTN6BQF7CP5IK3MJUAR4HV"
	DECIMAL = 32
	PAD     = "G"
	LEN     = 6
)

// 计算Json唯一ID
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

// 判断Json是否过期或被禁用
func IsExpiredOrForbidden(ExpirationTime time.Time, status uint) (err error) {
	//x := "2018-12-27 18:44:55"
	fmt.Println(ExpirationTime)

	nowTime := time.Now().Unix()

	loc, _ := time.LoadLocation("Asia/Shanghai")                                                            //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", ExpirationTime.Format("2006-01-02 15:04:05"), loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())

	fmt.Println(nowTime)

	if tt.Unix() < nowTime {
		return errors.New("Expired.")
	}
	if status != 0 {
		return errors.New("The Json has been disabled.")
	}
	return nil
}

// 文字转json
func String2Json(Str string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(Str), &dat)
	if err != nil {
		return nil, err
	} else {
		return dat, err
	}
}
