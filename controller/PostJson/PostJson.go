package PostJson

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/repository"
	"net/http"
	"strconv"

	//"jsonhutapi/repository"
)

type InputJson struct {
	Json string `json:"json" binding:"required,json"`
	Time string `json:"time" binding:"required,number"`
}

type ReturnData struct {
	Id string `json:"id"`
}

type ReturnJson struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ReturnData `json:"data"`
}

type ReturnJsonWithoutData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func PostJson(ctx *gin.Context) {
	//fmt.Println(ctx.Request.Header.Get("Origin")) //请求头部
	json := InputJson{}
	err := ctx.ShouldBind(&json)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Parameter error",
		})
		return
	}

	intNum, _ := strconv.Atoi(json.Time)
	uid, err := repository.CreateJson(json.Json, intNum)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Internal error",
		})
		return
	}
	//fmt.Println(uid)
	jsonId := repository.UpdateJsonID(uid)

	ctx.JSON(http.StatusCreated, ReturnJson{
		Code: 200,
		Msg:  "Success",
		Data: ReturnData{Id: jsonId},
	})
}
