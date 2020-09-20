package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
)

type GameHandler interface {
	HandleFinish() http.HandlerFunc
}

type gameHandler struct {
	gameUsecase usecase.GameUsecase
}

// NewGameHandler GameHandlerを生成
func NewGameHandler(gu usecase.GameUsecase) GameHandler {
	return &gameHandler{gameUsecase: gu}
}

// HandleFinish ユーザ情報更新処理
func (gh *gameHandler) HandleFinish() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type finishRequest struct {
			Score int32 `json:"score"`
			Stage int32 `json:"stage"`
		}

		type finishResponse struct {
			Coin int32 `json:"coin"`
		}

		// Contextから認証済みのユーザーを取得
		ctx := request.Context()
		user := dcontext.GetUserFromContext(ctx)
		if user == nil {
			log.Println("user is empty")
			response.BadRequest(writer, "user is empty")
			return
		}

		// リクエストBodyから、ゲームスコアscoreを取得
		var requestBody finishRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, `request parameter is empty`)
			return
		}

		// ゲーム終了時にユーザー情報を更新するユースケースを呼び出し
		getCoin, err := gh.gameUsecase.GameFinish(user, requestBody.Stage, requestBody.Score)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		response.Success(writer, &finishResponse{
			Coin: getCoin,
		})
	}
}
