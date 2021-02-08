package GetJsonDetails

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/repository"
	"net/http"
)

type ReturnData struct {
	JsonBody map[string]interface{} `json:"json_body"`
	Url      string                 `json:"url"`
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

func GetDetails(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, ReturnJson{
		Code: 200,
		Msg:  "Success",
		Data: ReturnData{
			JsonBody: dat,
			Url:      "https://api.jsonhut.com/bins/" + id,
		},
	})
}
