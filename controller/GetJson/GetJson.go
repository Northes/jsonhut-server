package GetJson

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/repository"
	"net/http"
)

type ReturnJsonWithoutData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func GetJson(ctx *gin.Context) {
	id := ctx.Param("id")

	resultData,err := repository.QueryJsonBodyByJsonID(id)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK,ReturnJsonWithoutData{
			Code: 400,
			Msg:  "Record not found",
		})
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(resultData.JsonBody), &dat); err != nil {
		ctx.JSON(http.StatusOK,ReturnJsonWithoutData{
			Code: 400,
			Msg:  "Unsupported data type",
		})
		return
	}
	//fmt.Println(ctx.Request.Header.Get("Origin")) //请求头部
	ctx.JSON(http.StatusOK, dat)
}
