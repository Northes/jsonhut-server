package logic

import (
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
