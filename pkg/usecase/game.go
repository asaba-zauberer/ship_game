package usecase

import (
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
)

type GameUsecase interface {
	GameFinish(record *model.User, score, stage int32) (getCoin int32, err error)
}

type gameUsecase struct {
	userRepo repository.UserRepo
}

// NewGameUsecase GameUsecaseを生成
func NewGameUsecase(ur repository.UserRepo) GameUsecase {
	return &gameUsecase{userRepo: ur}
}

// GameFinish ゲーム終了時のユースケース
// TODO リニューアル
func (uu gameUsecase) GameFinish(record *model.User, score, stage int32) (int32, error) {

	// 獲得コイン
	//getCoin := score
	//coin := record.Coin + getCoin
	/*
		// ハイスコア更新を確認
		var highScore int32
		if record.HighScore >= score {
			highScore = record.HighScore
		} else {
			highScore = score
		}
	*/
	var coin int32 = 100000
	if err := uu.userRepo.Update(record, record.Name, coin, record.Stage); err != nil {
		return 0, err
	}

	// TODO スコアテーブル更新

	//return getCoin, nil

	return 100000, nil
}
