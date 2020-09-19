package usecase

import (
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
)

type RankingUsecase interface {
	GetRanking(start int) (model.Users, error)
}

type rankingUsecase struct {
	userRepo repository.UserRepo
}

// NewRankingUsecase RankingUsecaseを生成
func NewRankingUsecase(ur repository.UserRepo) RankingUsecase {
	return &rankingUsecase{userRepo: ur}
}

// GetRanking ランキング表示のユースケース
// TODO リニューアル
func (ru rankingUsecase) GetRanking(start int) (model.Users, error) {
	//users, err := ru.userRepo.SelectByHighScore(start)
	/*
		if err != nil {
			return nil, err
		}
		return users, nil
	*/
	return nil, nil
}
