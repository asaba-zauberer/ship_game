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
		start, err := strconv.Atoi(request.URL.Query().Get("start"))
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "query 'start' is empty")
			return
		}
		if start < 1 || start > constant.NumberOfStage {
			log.Printf("query 'start' is invalid. start=%d", start)
			response.BadRequest(writer, "query 'start' is invalid")
			return
		}

		// ユーザー情報の取得を行うユースケースを呼び出し
		users, err := rh.rankingUsecase.GetRanking(start)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		rankingList := make([]RankingListItem, len(users))
		/*
			for i, user := range users {
				rankingList[i] = RankingListItem{
					UserID:   user.ID,
					UserName: user.Name,
					Rank:     int32(start + i),
					Score:    user.HighScore,
				}
			}
		*/
		response.Success(writer, &RankingListResponse{
			Ranks: rankingList,
		})
	}
}

type RankingListItem struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Rank     int32  `json:"rank"`
	Score    int32  `json:"score"`
}

type RankingListResponse struct {
	Ranks []RankingListItem `json:"ranks"`
}
