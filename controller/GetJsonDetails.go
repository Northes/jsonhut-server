package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/models"
	"net/http"
)

func GetJsonDetails(ctx *gin.Context) {
	jsonID := ctx.Param("id")

	resultData, err := models.QueryJsonBodyByJsonID(jsonID)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
			Code: 404,
			Msg:  "Record not found",
		})
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(resultData.JsonBody), &dat); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Unsupported data type",
		})
		return
	}

	//loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	//tt, _ := time.ParseInLocation("2006-01-02 15:04:05", resultData.ExpirationTime.Format("2006-01-02 15:04:05"), loc)
	eTime := resultData.ExpirationTime.Format("2006-01-02 15:04:05")
	cTime := resultData.CreatedAt.Format("2006-01-02 15:04:05")
	uTime := resultData.UpdatedAt.Format("2006-01-02 15:04:05")

	ctx.JSON(http.StatusOK, models.DetailsReturnJson{
		Code: 200,
		Msg:  "Success",
		Data: models.DetailsReturnData{
			//JsonBody:       dat,
			Url:            "https://api.jsonhut.com/bins/" + jsonID,
			Count:          resultData.CallCount,
			ExpirationTime: eTime,
			CreatedAt:      cTime,
			UpdatedAt:      uTime,
		},
	})
}
