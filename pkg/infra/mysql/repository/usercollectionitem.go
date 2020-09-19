package repository

import (
	"database/sql"
	"log"

	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
)

type UserCollectionItemRepo interface {
	SelectByUserID(userID string) (model.UserCollectionItems, error)
	InsertItems(userCollectionItems model.UserCollectionItems) error
}

type userCollectionItemRepo struct {
	mysql.SqlHandler
}

// NewUserCollectionItemRepo repository/usercollectionitem.goで用意したインタフェースの実装
func NewUserCollectionItemRepo(sqlHandler mysql.SqlHandler) repository.UserCollectionItemRepo {
	userCollectionItemRepo := userCollectionItemRepo{sqlHandler}
	return &userCollectionItemRepo
}

// SelectByUserID userIDを条件にレコードを取得
func (ucr *userCollectionItemRepo) SelectByUserID(userID string) (model.UserCollectionItems, error) {
	rows, err := ucr.SqlHandler.Conn.Query("SELECT * FROM user_collection_item WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return convertToUserCollectionItems(rows)
}

// InsertItems 複数のアイテムを保存
func (ucr *userCollectionItemRepo) InsertItems(userCollectionItems model.UserCollectionItems) error {
	query, queryParam := createInsertItemsQuery(userCollectionItems)
	stmt, err := ucr.SqlHandler.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(queryParam...)
	return err
}

// CreateInsertItemsQuery 複数のアイテムを保存するクエリを作成
func createInsertItemsQuery(uc model.UserCollectionItems) (string, []interface{}) {
	query := "INSERT IGNORE INTO user_collection_item(user_id, collection_item_id) VALUES"
	var queryParam []interface{}

	for n, u := range uc {
		query += "(?, ?)"
		if n+1 != len(uc) {
			query += ","
		}
		queryParam = append(queryParam, u.UserID, u.CollectionItemID)

	}
	return query, queryParam
}

// convertToUserCollectionItems rowsデータをCollectionsデータへ変換する
func convertToUserCollectionItems(rows *sql.Rows) (model.UserCollectionItems, error) {
	var userCollectionItems model.UserCollectionItems

	for rows.Next() {
		userCollectionItem := model.UserCollectionItem{}
		if err := rows.Scan(&userCollectionItem.UserID, &userCollectionItem.CollectionItemID); err != nil {
			log.Println(err)
			return nil, err
		}
		userCollectionItems = append(userCollectionItems, &userCollectionItem)
	}

	return userCollectionItems, nil
}
