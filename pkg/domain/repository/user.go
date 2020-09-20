package repository

import (
	"20dojo-online/pkg/domain/model"
)

// Repository Userモデルの永続化、依存関係を守るため実装はinfra
type UserRepo interface {
	Create(id, authToken, name string, coin, stage int32) error
	SelectByAuthToken(authToken string) (*model.User, error)
	SelectByID(id []string) (model.Users, error)
	Update(record *model.User, name string, coin, stage int32) error
	Lock(*model.User) error
}
