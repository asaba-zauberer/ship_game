package handler

import (
	"log"
	"net/http"

	"20dojo-online/pkg/interfaces/dcontext"
	"20dojo-online/pkg/interfaces/response"
	"20dojo-online/pkg/usecase"
)

type CollectionHandler interface {
	HandleList() http.HandlerFunc
}

type collectionHandler struct {
	collectionUsecase usecase.CollectionUsecase
}

// NewCollectionHandler コレクションに関するHandlerを生成
func NewCollectionHandler(cu usecase.CollectionUsecase) CollectionHandler {
	return &collectionHandler{collectionUsecase: cu}
}

// HandleList コレクションリスト取得処理
func (ch *collectionHandler) HandleList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type collectionListItem struct {
			CollectionID string `json:"collectionID"`
			Name         string `json:"name"`
			Rarity       int32  `json:"rarity"`
			HasItem      bool   `json:"hasItem"`
		}

		type collectionListResponse struct {
			Collections []collectionListItem `json:"collections"`
		}

		// Contextから認証済みのユーザーを取得
		ctx := request.Context()
		user := dcontext.GetUserFromContext(ctx)
		if user == nil {
			log.Println("user is empty")
			response.BadRequest(writer, "user is empty")
			return
		}

		// アイテム一覧とその所有情報を取得するユースケースを呼び出し
		collectionItems, err := ch.collectionUsecase.GetCollection(user.ID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		collectionResult := make([]collectionListItem, len(collectionItems))
		for i, collectionItems := range collectionItems {
			collectionResult[i] = collectionListItem{
				CollectionID: collectionItems.ID,
				Name:         collectionItems.Name,
				Rarity:       collectionItems.Rarity,
				HasItem:      collectionItems.HasItem,
			}
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, &collectionListResponse{
			Collections: collectionResult,
		})
	}
}
