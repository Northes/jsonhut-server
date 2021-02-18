package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/dao"
	"jsonhutapi/logic"
	"jsonhutapi/models"
	"net/http"
)

func GetJson(ctx *gin.Context) {
	jsonID := ctx.Param("id")
	from := ctx.Query("from")
	// 尝试从Redis中获取缓存数据
	redisResult, err := dao.RedisGetData(jsonID)
	if err == nil {
		fmt.Printf("From Redis : %s\n", redisResult)
		json, _ := logic.String2Json(redisResult)
		ctx.JSON(http.StatusOK, json)
		// 调用时增加次数（利用Gorm自动更新调用时间），从详情页访问时不计数
		if from != "details" {
			models.UpdateJsonCallCount(jsonID)
		}
		return
	}

	// 从MySQl中获取数据
	mysqlResult, err := models.QueryJsonBodyByJsonID(jsonID)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
			Code: 404,
			Msg:  "Record not found",
		})
		return
	}

	// 判断是否过期或禁用
	if err = logic.IsExpiredOrForbidden(mysqlResult.ExpirationTime, mysqlResult.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ReturnJsonWithoutData{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	// Str2Json
	dat, err := logic.String2Json(mysqlResult.JsonBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Unsupported data type",
		})
		return
	}

	// 调用时增加次数
	if from != "details" {
		models.UpdateJsonCallCount(jsonID)
	}
	// 写入Redis缓存
	dao.RedisSetData(mysqlResult.JsonId, mysqlResult.JsonBody)
	// 设置Redis过期时间
	dao.RedisSetExpirationTime(mysqlResult.JsonId, -1)
	ctx.JSON(http.StatusOK, dat)
}
