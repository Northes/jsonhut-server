package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/logic"
	"jsonhutapi/models"
	"net/http"
)

func GetJson(ctx *gin.Context) {
	id := ctx.Param("id")
	// 取到数据
	resultData, err := models.QueryJsonBodyByJsonID(id)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
			Code: 404,
			Msg:  "Record not found",
		})
		return
	}
	// 判断是否过期或禁用
	if err = logic.IsExpiredOrForbidden(resultData.ExpirationTime, resultData.Status); err != nil {
		ctx.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
			Code: 404,
			Msg:  err.Error(),
		})
		return
	}

	// 反序列化
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(resultData.JsonBody), &dat); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Unsupported data type",
		})
		return
	}
	//fmt.Println(ctx.Request.Header.Get("Origin")) //请求头部
	ctx.JSON(http.StatusOK, dat)
}
