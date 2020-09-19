package repository

import (
	"20dojo-online/pkg/domain/model"
)

// Repository GachaProbabilityモデルの永続化、依存関係を守るため実装はinfra
type GachaProbabilityRepo interface {
	SelectAll() (model.GachaProbabilities, error)
}
