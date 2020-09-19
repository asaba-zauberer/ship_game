package model

// UserCollectionItem アイテム所持情報のモデル
type UserCollectionItem struct {
	UserID           string
	CollectionItemID string
}

// UserCollectionItems アイテム所持情報のスライス
type UserCollectionItems []*UserCollectionItem
