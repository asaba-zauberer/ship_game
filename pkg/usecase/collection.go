package usecase

import (
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
)

type CollectionUsecase interface {
	GetCollection(userID string) (model.CollectionItems, map[string]struct{}, error)
}

type collectionUsecase struct {
	collectionItemRepo     repository.CollectionItemRepo
	userCollectionItemRepo repository.UserCollectionItemRepo
}

// NewCollectionUsecase CollectionUsecaseを生成
func NewCollectionUsecase(cr repository.CollectionItemRepo, ucr repository.UserCollectionItemRepo) CollectionUsecase {
	return &collectionUsecase{
		collectionItemRepo:     cr,
		userCollectionItemRepo: ucr,
	}
}

// GetCollection userIDを条件に全てのアイテムとその所持情報を取得
func (cu collectionUsecase) GetCollection(userID string) (model.CollectionItems, map[string]struct{}, error) {
	collectionItems, err := cu.collectionItemRepo.SelectAll()
	if err != nil {
		return nil, nil, err
	}

	userCollectionItems, err := cu.userCollectionItemRepo.SelectByUserID(userID)
	if err != nil {
		return nil, nil, err
	}

	hasCollectionItem := make(map[string]struct{}, len(userCollectionItems))
	for _, userCollectionItem := range userCollectionItems {
		id := userCollectionItem.CollectionItemID
		hasCollectionItem[id] = struct{}{}
	}

	return collectionItems, hasCollectionItem, nil
}
