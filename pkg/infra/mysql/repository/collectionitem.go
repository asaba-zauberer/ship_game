package repository

import (
	"database/sql"
	"log"

	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
)

type CollectionItemRepo interface {
	SelectAll() (model.CollectionItems, error)
	SelectByID(ID []string) (model.CollectionItems, error)
}

type collectionItemRepo struct {
	mysql.SqlHandler
}

// NewCollectionItemRepo repository/collectionitem.goで用意したインタフェースの実装
func NewCollectionItemRepo(sqlHandler mysql.SqlHandler) repository.CollectionItemRepo {
	collectionItemRepo := collectionItemRepo{sqlHandler}
	return &collectionItemRepo
}

// SelectAll 全てのアイテムのレコードを取得
func (cr *collectionItemRepo) SelectAll() (model.CollectionItems, error) {
	rows, err := cr.SqlHandler.Conn.Query("SELECT * FROM collection_item")
	if err != nil {
		return nil, err
	}
	return convertToCollectionItems(rows)
}

// SelectByID IDを条件にアイテムレコードを取得
func (cr *collectionItemRepo) SelectByID(ID []string) (model.CollectionItems, error) {
	query, queryParam := createSelectByIDQuery(ID)
	rows, err := cr.SqlHandler.Conn.Query(query, queryParam...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return convertToCollectionItems(rows)
}

// createSelectByIDQuery IDを条件にレコードを取得するクエリを作成
func createSelectByIDQuery(ID []string) (string, []interface{}) {
	query := "SELECT * FROM collection_item WHERE id IN ("
	var queryParam []interface{}

	for n, i := range ID {
		query += "?"
		if n+1 != len(ID) {
			query += ","
		} else {
			query += ")"
		}
		queryParam = append(queryParam, i)
	}

	return query, queryParam
}

// convertToCollection rowデータをcollectionデータに変換
func convertToCollectionItems(rows *sql.Rows) (model.CollectionItems, error) {
	var collectionItems model.CollectionItems

	for rows.Next() {
		collectionItem := model.CollectionItem{}
		if err := rows.Scan(&collectionItem.ID, &collectionItem.Name, &collectionItem.Rarity); err != nil {
			log.Println(err)
			return nil, err
		}
		collectionItems = append(collectionItems, &collectionItem)
	}

	return collectionItems, nil
}
