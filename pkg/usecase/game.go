package usecase

import (
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
)

type GameUsecase interface {
	GameFinish(record *model.User, score, stage int32) (getCoin int32, err error)
}

type gameUsecase struct {
	userRepo      repository.UserRepo
	userScoreRepo repository.UserScoreRepo
}

// NewGameUsecase GameUsecaseを生成
func NewGameUsecase(ur repository.UserRepo, usr repository.UserScoreRepo) GameUsecase {
	return &gameUsecase{
		userRepo:      ur,
		userScoreRepo: usr,
	}
}

// GameFinish ゲーム終了時のユースケース
func (uu gameUsecase) GameFinish(user *model.User, stage, score int32) (int32, error) {
	var coin int32

	// 初回クリア時のみスコア多めに付与
	if user.Stage < stage {
		coin = 3

		// スコア登録
		if err := uu.userScoreRepo.InsertScore(user.ID, stage, score); err != nil {
			return 0, err
		}

		// ユーザー到達ステージ更新
		if err := uu.userRepo.Update(user, user.Name, user.Coin+coin, stage); err != nil {
			return 0, err
		}
	} else {
		coin = 1

		// 過去のスコアを取得
		userScore, err := uu.userScoreRepo.GetScoreByIDAndStage(user.ID, stage)
		if err != nil {
			return 0, err
		}
		// ハイスコアであれば更新
		if userScore.Score < score {
			if err := uu.userScoreRepo.UpdateScore(user.ID, stage, score); err != nil {
				return 0, err
			}
		}
	}

	return coin, nil
}
