package model

// CollectionItem アイテムを表すモデル
type CollectionItem struct {
	ID     string
	Name   string
	Rarity int32
}

// CollectionItems アイテムのスライス
type CollectionItems []*CollectionItem
