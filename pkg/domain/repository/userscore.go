package repository

import (
	"20dojo-online/pkg/domain/model"
)

// Repository Userモデルの永続化、依存関係を守るため実装はinfra
type UserScoreRepo interface {
	GetScoreByIDAndStage(id string, stage int32) (*model.UserScore, error)
	GetScoresByID(id string) (model.UserScores, error)
	InsertScore(id string, stage, score int32) error
	UpdateScore(id string, stage, score int32) error
	GetRankingByStage(stage int) (model.UserScores, error)
}
