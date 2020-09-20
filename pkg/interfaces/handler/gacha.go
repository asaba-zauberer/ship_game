package handler

import (
	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type GachaHandler interface {
	HandleDraw() http.HandlerFunc
}
type gachaHandler struct {
	gachaUsecase usecase.GachaUsecase
}

// NewHandler GachaHandlerを生成
func NewGachaHandler(gu usecase.GachaUsecase) GachaHandler {
	return &gachaHandler{gachaUsecase: gu}
}

// HandleDraw ガチャ処理
func (gh gachaHandler) HandleDraw() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type drawRequest struct {
			Times int32 `json:"times"`
		}

		type resultItem struct {
			CollectionID string `json:"collectionID"`
			Name         string `json:"name"`
			Rarity       int32  `json:"rarity"`
			IsNew        bool   `json:"isNew"`
		}

		type resultResponse struct {
			Results []resultItem `json:"results"`
		}

		// リクエストBodyからガチャ回数timesを取得
		var requestBody drawRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, `query "times" is empty`)
			return
		}

		// Contextから認証済みのユーザーを取得
		ctx := request.Context()
		user := dcontext.GetUserFromContext(ctx)
		if user == nil {
			log.Println("user is empty")
			response.BadRequest(writer, "user is empty")
			return
		}

		// ガチャ結果を取得
		gachaResponse, err := gh.gachaUsecase.DrawGacha(user, requestBody.Times)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// JSONに整形
		gachaResult := make([]resultItem, len(gachaResponse))
		for i, gachaResponseItem := range gachaResponse {
			gachaResult[i] = resultItem{
				CollectionID: gachaResponseItem.ID,
				Name:         gachaResponseItem.Name,
				Rarity:       gachaResponseItem.Rarity,
				IsNew:        gachaResponseItem.IsNew,
			}
		}

		//response.Success(writer, gachaResult)
		response.Success(writer, &resultResponse{
			Results: gachaResult,
		})
	}
}
