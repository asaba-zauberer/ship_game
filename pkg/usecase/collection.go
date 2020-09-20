package usecase

import (
	"20dojo-online/pkg/domain/repository"
)

type CollectionUsecase interface {
	GetCollection(userID string) ([]*collectionResult, error)
}

type collectionUsecase struct {
	collectionItemRepo     repository.CollectionItemRepo
	userCollectionItemRepo repository.UserCollectionItemRepo
}

type collectionResult struct {
	ID      string
	Name    string
	Rarity  int32
	HasItem bool
}

// NewCollectionUsecase CollectionUsecaseを生成
func NewCollectionUsecase(cr repository.CollectionItemRepo, ucr repository.UserCollectionItemRepo) CollectionUsecase {
	return &collectionUsecase{
		collectionItemRepo:     cr,
		userCollectionItemRepo: ucr,
	}
}

// GetCollection userIDを条件に全てのアイテムとその所持情報を取得
func (cu collectionUsecase) GetCollection(userID string) ([]*collectionResult, error) {
	collectionItems, err := cu.collectionItemRepo.SelectAll()
	if err != nil {
		return nil, err
	}

	userCollectionItems, err := cu.userCollectionItemRepo.SelectByUserID(userID)
	if err != nil {
		return nil, err
	}

	hasCollectionItem := make(map[string]struct{}, len(userCollectionItems))
	for _, userCollectionItem := range userCollectionItems {
		id := userCollectionItem.CollectionItemID
		hasCollectionItem[id] = struct{}{}
	}

	// レスポンスのhasItemを作成
	collectionList := make([]*collectionResult, len(collectionItems))
	for i, collectionItem := range collectionItems {
		_, ok := hasCollectionItem[collectionItem.ID]
		hasItem := ok

		collectionList[i] = &collectionResult{
			ID:      collectionItem.ID,
			Name:    collectionItem.Name,
			Rarity:  collectionItem.Rarity,
			HasItem: hasItem,
		}
	}

	return collectionList, nil
}
