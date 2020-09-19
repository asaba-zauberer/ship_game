package model

// GachaProbability アイテム排出率を表すモデル
type GachaProbability struct {
	CollectionItemID string
	Ratio            int32
}

// GachaProbabilities アイテムの排出率のスライス
type GachaProbabilities []*GachaProbability
