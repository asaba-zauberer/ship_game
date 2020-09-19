package repository

import (
	"20dojo-online/pkg/domain/model"
)

// Repository UserCollectionItemモデルの永続化、依存関係を守るため実装はinfra
type UserCollectionItemRepo interface {
	SelectByUserID(userID string) (model.UserCollectionItems, error)
	InsertItems(model.UserCollectionItems) error
}
