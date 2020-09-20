package handler

import (
	"log"
	"net/http"
	"strconv"

	"20dojo-online/constant"
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
		// URLクエリパラメータを取得
		stage, err := strconv.Atoi(request.URL.Query().Get("stage"))
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "query 'stage' is empty")
			return
		}
		if stage < 1 || stage > constant.NumberOfStage {
			log.Printf("query 'stage' is invalid. start=%d", stage)
			response.BadRequest(writer, "query 'stage' is invalid")
			return
		}

		// ユーザー情報の取得を行うユースケースを呼び出し
		rankingUserList, err := rh.rankingUsecase.GetRanking(stage)
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
