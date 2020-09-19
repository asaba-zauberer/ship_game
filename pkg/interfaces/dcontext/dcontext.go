package dcontext

import (
	"20dojo-online/pkg/domain/model"
	"context"
)

type key string

const (
	userIDKey key = "userID"
)

// SetUser ContextへユーザIDを保存する
func SetUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userIDKey, user)
}

// GetUserFromContext Contextからユーザーを取得する
func GetUserFromContext(ctx context.Context) *model.User {
	var user model.User
	if ctx.Value(userIDKey) == nil {
		return nil
	}
	user = ctx.Value(userIDKey).(model.User)
	return &user
}
