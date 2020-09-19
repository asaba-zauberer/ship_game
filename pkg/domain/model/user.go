package model

// User ユーザーのモデル
type User struct {
	ID        string
	AuthToken string
	Name      string
	Coin      int32
	Stage     int32
}

// Users Userのスライス
type Users []*User
