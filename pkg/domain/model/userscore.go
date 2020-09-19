package model

// UserScore ユーザースコアを表すモデル
type UserScore struct {
	ID    string
	Stage int32
	Score int32
}

// GachaProbabilities ユーザースコアのスライス
type UserScores []*UserScore
