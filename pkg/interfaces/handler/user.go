package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
)

type UserHandler interface {
	HandleCreate() http.HandlerFunc
	HandleGet() http.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewHandle UserHandlerを生成
func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

// HandleCreate ユーザー情報の作成
func (uh *userHandler) HandleCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type createRequest struct {
			Name string `json:"name"`
		}

		type createResponse struct {
			Token string `json:"token"`
		}

		// リクエストBodyから、新規UserのNameを取得
		var requestBody createRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, `query "name" empty`)
			return
		}

		authToken, err := uh.userUsecase.Create(requestBody.Name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 生成したユーザーから認証トークンを返却
		response.Success(writer, &createResponse{Token: authToken})
	}
}

// HandleGet ユーザ情報の取得
func (uh *userHandler) HandleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type highScore struct {
			Stage     int32 `json:"stage"`
			HighScore int32 `json:"high_score"`
		}

		type getResponse struct {
			ID         string      `json:"id"`
			Name       string      `json:"name"`
			HighScores []highScore `json:"high_scores"`
			Coin       int32       `json:"coin"`
			Stage      int32       `json:"stage"`
		}

		// Contextから認証済みのユーザーを取得
		ctx := request.Context()
		user := dcontext.GetUserFromContext(ctx)
		if user == nil {
			log.Println("user is empty")
			response.BadRequest(writer, "user is empty")
			return
		}

		// スコア情報の取得を行うユースケースを呼び出し
		userScores, err := uh.userUsecase.GetHighScoresByID(user.ID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンスのhighScoresを作成
		highScoreList := make([]highScore, len(userScores))
		for i, scoreData := range userScores {
			highScoreList[i] = highScore{
				Stage:     scoreData.Stage,
				HighScore: scoreData.Score,
			}
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, &getResponse{
			ID:         user.ID,
			Name:       user.Name,
			HighScores: highScoreList,
			Coin:       user.Coin,
			Stage:      user.Stage,
		})
	}
}
