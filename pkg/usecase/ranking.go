package usecase

import (
	"20dojo-online/pkg/domain/repository"
)

type RankingUsecase interface {
	GetRanking(stage int32) ([]*rankingResponse, error)
}

type rankingUsecase struct {
	userRepo      repository.UserRepo
	userScoreRepo repository.UserScoreRepo
}

type rankingResponse struct {
	UserID   string
	UserName string
	Rank     int32
	Score    int32
}

// NewRankingUsecase RankingUsecaseを生成
func NewRankingUsecase(ur repository.UserRepo, usr repository.UserScoreRepo) RankingUsecase {
	return &rankingUsecase{
		userRepo:      ur,
		userScoreRepo: usr,
	}
}

// GetRanking ランキング表示のユースケース
func (ru rankingUsecase) GetRanking(stage int32) ([]*rankingResponse, error) {
	// ステージを条件にランキングを取得
	userScores, err := ru.userScoreRepo.GetRankingByStage(stage)
	if err != nil {
		return nil, err
	}

	// ユーザー情報を取得
	userIDList := make([]string, len(userScores))
	for i, score := range userScores {
		userIDList[i] = score.ID
	}

	users, err := ru.userRepo.SelectByID(userIDList)
	if err != nil {
		return nil, err
	}

	// マップに変換
	usersMap := make(map[string]string)
	for _, data := range users {
		usersMap[data.ID] = data.Name
	}

	// 戻り値の作成
	rankingList := make([]*rankingResponse, len(userScores))
	var previousScore int32
	var rank int32
	for i, scoreData := range userScores {
		if scoreData.Score != previousScore {
			rank = int32(i + 1)
			previousScore = scoreData.Score
		}
		rankingList[i] = &rankingResponse{
			UserID:   scoreData.ID,
			UserName: usersMap[scoreData.ID],
			Rank:     rank,
			Score:    scoreData.Score,
		}
	}

	return rankingList, nil
}
