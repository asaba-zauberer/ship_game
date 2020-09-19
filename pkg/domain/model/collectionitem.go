package model

// CollectionItem アイテムを表すモデル
type CollectionItem struct {
	ID   string
	Name string
}

// CollectionItems アイテムのスライス
type CollectionItems []*CollectionItem
