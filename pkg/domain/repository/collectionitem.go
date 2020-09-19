package repository

import (
	"20dojo-online/pkg/domain/model"
)

// Repository CollectionItemモデルの永続化、依存関係を守るため実装はinfra
type CollectionItemRepo interface {
	SelectAll() (model.CollectionItems, error)
	SelectByID(ID []string) (model.CollectionItems, error)
}
