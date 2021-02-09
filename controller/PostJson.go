package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/models"
	"net/http"
	"strconv"
)

func PostJson(ctx *gin.Context) {
	//fmt.Println(ctx.Request.Header.Get("Origin")) //请求头部
	json := models.PostInputJson{}
	err := ctx.ShouldBind(&json)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, models.ReturnJsonWithoutData{
			Code: 400,
			Msg:  "Parameter error",
		})
		return
	}

	intNum, _ := strconv.Atoi(json.Day)
	uid, err := models.CreateAJson(json.Json, intNum)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Internal error",
		})
		return
	}
	//fmt.Println(uid)
	jsonId := models.UpdateJsonID(uid)

	ctx.JSON(http.StatusCreated, models.PostReturnJson{
		Code: 201,
		Msg:  "Success",
		Data: models.PostReturnData{Id: jsonId},
	})
}
