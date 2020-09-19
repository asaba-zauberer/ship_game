package handler

import (
	"net/http"

	"20dojo-online/constant"
	"20dojo-online/pkg/interfaces/response"
)

// HandleGet ゲーム設定情報取得処理
func HandleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type getResponse struct {
			GachaCoinConsumption int32 `json:"gachaCoinConsumption"`
		}
		response.Success(writer, &getResponse{
			GachaCoinConsumption: constant.GachaCoinConsumption,
		})
	}
}
