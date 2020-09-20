package server

import (
	"20dojo-online/pkg/injector"
	"20dojo-online/pkg/interfaces/handler"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Serve(addr string) {
	// 依存性の解決
	middleware := injector.InjectMiddleware()
	userHandler := injector.InjectUserHandler()
	gameHandler := injector.InjectGameHandler()
	rankingHandler := injector.InjectRankingHandler()
	collectionHandler := injector.InjectCollectionHandler()
	gachaHandler := injector.InjectGachaHandler()

	// ルーティング
	http.HandleFunc("/setting/get", get(handler.HandleGet()))

	http.HandleFunc("/user/create", post(userHandler.HandleCreate()))
	http.HandleFunc("/user/get", get(middleware.Authenticate(userHandler.HandleGet())))

	http.HandleFunc("/game/finish", post(middleware.Authenticate(gameHandler.HandleFinish())))

	http.HandleFunc("/ranking/list", post(middleware.Authenticate(rankingHandler.HandleList())))

	http.HandleFunc("/collection/list", get(middleware.Authenticate(collectionHandler.HandleList())))

	http.HandleFunc("/gacha/draw", post(middleware.Authenticate(gachaHandler.HandleDraw())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
