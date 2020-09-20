package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/constant"
	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
)

type RankingHandler interface {
	HandleList() http.HandlerFunc
}

type rankingHandler struct {
	rankingUsecase usecase.RankingUsecase
}

// NewRankingHandler RankingHandlerの作成
func NewRankingHandler(ru usecase.RankingUsecase) RankingHandler {
	return &rankingHandler{rankingUsecase: ru}
}

// HandleList ランキング取得処理
func (rh *rankingHandler) HandleList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type rankingRequest struct {
			Stage int32 `json:"stage"`
		}

		// Contextから認証済みのユーザーを取得
		ctx := request.Context()
		user := dcontext.GetUserFromContext(ctx)
		if user == nil {
			log.Println("user is empty")
			response.BadRequest(writer, "user is empty")
			return
		}

		// リクエストBodyから、ステージを取得
		var requestBody rankingRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, `request parameter is empty`)
			return
		}
		if requestBody.Stage < 1 || requestBody.Stage > constant.NumberOfStage {
			log.Printf("query 'stage' is invalid. start=%d", requestBody.Stage)
			response.BadRequest(writer, "query 'stage' is invalid")
			return
		}

		// ユーザー情報の取得を行うユースケースを呼び出し
		rankingUserList, err := rh.rankingUsecase.GetRanking(requestBody.Stage)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// JSONに整形
		rankingResult := make([]rankingListItem, len(rankingUserList))
		for i, rankUserData := range rankingUserList {
			rankingResult[i] = rankingListItem{
				UserID:   rankUserData.UserID,
				UserName: rankUserData.UserName,
				Rank:     int32(i) + 1,
				Score:    rankUserData.Score,
			}
		}

		response.Success(writer, &rankingListResponse{
			Ranks: rankingResult,
		})
	}
}

type rankingListItem struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Rank     int32  `json:"rank"`
	Score    int32  `json:"score"`
}

type rankingListResponse struct {
	Ranks []rankingListItem `json:"ranks"`
}
