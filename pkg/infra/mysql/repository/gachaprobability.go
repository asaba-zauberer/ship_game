package repository

import (
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
	"database/sql"
	"log"
)

type GachaProbabilityRepo interface {
	SelectAll() (model.GachaProbabilities, error)
}

type gachaProbabilityRepo struct {
	mysql.SqlHandler
}

// NewGachaProbabilityRepo repository/gachaprobability.goで用意したインタフェースの実装
func NewGachaProbabilityRepo(sqlHandler mysql.SqlHandler) repository.GachaProbabilityRepo {
	gachaProbabilityRepo := gachaProbabilityRepo{sqlHandler}
	return &gachaProbabilityRepo
}

// SelectAll 全てのガチャアイテム排出率を取得
func (gr *gachaProbabilityRepo) SelectAll() (model.GachaProbabilities, error) {
	rows, err := gr.SqlHandler.Conn.Query("SELECT * FROM gacha_probability")
	if err != nil {
		return nil, err
	}
	return convertToGachaItems(rows)
}

// convertToGachas rowsデータをGachaProbabilitiesデータへ変換、Ratioを累積値に変更
func convertToGachaItems(rows *sql.Rows) (model.GachaProbabilities, error) {
	var gachaProbabilities model.GachaProbabilities
	var sumOfProb int32

	for rows.Next() {
		gachaProbability := model.GachaProbability{}
		if err := rows.Scan(&gachaProbability.CollectionItemID, &gachaProbability.Ratio); err != nil {
			log.Println(err)
			return nil, err
		}
		sumOfProb += gachaProbability.Ratio
		gachaProbability.Ratio = sumOfProb
		gachaProbabilities = append(gachaProbabilities, &gachaProbability)
	}

	return gachaProbabilities, nil
}
