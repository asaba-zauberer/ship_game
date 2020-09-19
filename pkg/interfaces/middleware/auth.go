package middleware

import (
	"context"
	"log"
	"net/http"

	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
)

// Middleware middlewareのインターフェース
type Middleware interface {
	Authenticate(http.HandlerFunc) http.HandlerFunc
}

type middleware struct {
	usecase usecase.UserUsecase
}

// NewMiddleware userUseCaseと疎通
func NewMiddleware(uu usecase.UserUsecase) Middleware {
	return &middleware{
		usecase: uu,
	}
}

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func (m middleware) Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if token == "" {
			log.Println("x-token is empty")
			response.BadRequest(writer, "x-token is empty")
			return
		}

		// データベースから認証トークンに紐づくユーザの情報を取得
		user, err := m.usecase.SelectByAuthToken(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Invalid token")
			return
		}
		if user == nil {
			log.Printf("user not found. token=%s", token)
			response.BadRequest(writer, "Invalid token")
			return
		}
		// ユーザをContextへ保存して以降の処理に利用する
		ctx = dcontext.SetUser(ctx, *user)
		// 次の処理
		nextFunc(writer, request.WithContext(ctx))
	}
}
