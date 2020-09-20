package repository

import (
	"database/sql"
	"log"

	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
)

type UserScoreRepo interface {
	GetScoreByIDAndStage(id string, stage int32) (*model.UserScore, error)
	GetScoresByID(id string) (model.UserScores, error)
	InsertScore(id string, stage, score int32) error
	UpdateScore(id string, stage, score int32) error
	GetRankingByStage(stage int32) (model.UserScores, error)
}

type userScoreRepo struct {
	mysql.SqlHandler
}

// NewUserScoreRepo repository/usercollectionitem.goで用意したインタフェースの実装
func NewUserScoreRepo(sqlHandler mysql.SqlHandler) repository.UserScoreRepo {
	userScoreRepo := userScoreRepo{sqlHandler}
	return &userScoreRepo
}

// GetScoreByUserIDAndStage userID, stageを条件に1レコードを取得
func (usr *userScoreRepo) GetScoreByIDAndStage(userID string, stage int32) (*model.UserScore, error) {
	row := usr.SqlHandler.Conn.QueryRow("SELECT * FROM user_score WHERE id = ? AND stage = ?", userID, stage)
	return convertToUserScore(row)
}

// GetScoreByUserID userIDを条件に複数レコードを取得
func (usr *userScoreRepo) GetScoresByID(userID string) (model.UserScores, error) {
	rows, err := usr.SqlHandler.Conn.Query("SELECT * FROM user_score WHERE id = ?", userID)
	if err != nil {
		return nil, err
	}
	return convertToUserScores(rows)
}

// InsertScore スコアを登録
func (usr *userScoreRepo) InsertScore(userID string, stage, score int32) error {
	stmt, err := usr.SqlHandler.Conn.Prepare("INSERT INTO user_score (id, stage, score) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, stage, score)
	return err
}

// Update ユーザー情報の更新
func (usr *userScoreRepo) UpdateScore(userID string, stage, score int32) error {
	stmt, err := usr.SqlHandler.Conn.Prepare("UPDATE user_score SET score = ? WHERE id = ? AND stage = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(score, userID, stage)
	return err
}

// GetRankingByStage 指定ステージのランキングを取得
func (usr *userScoreRepo) GetRankingByStage(stage int32) (model.UserScores, error) {
	rows, err := usr.SqlHandler.Conn.Query("SELECT * FROM user_score WHERE stage = ? ORDER BY score DESC", stage)
	if err != nil {
		return nil, err
	}
	return convertToUserScores(rows)
}

// convertToUserScore rowデータをUserScoreデータへ変換
func convertToUserScore(row *sql.Row) (*model.UserScore, error) {
	userScore := model.UserScore{}
	if err := row.Scan(&userScore.ID, &userScore.Stage, &userScore.Score); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &userScore, nil
}

// convertToUserScores rowsデータをUserScoresデータへ変換する
func convertToUserScores(rows *sql.Rows) (model.UserScores, error) {
	var userScores model.UserScores

	for rows.Next() {
		userScore := model.UserScore{}
		if err := rows.Scan(&userScore.ID, &userScore.Stage, &userScore.Score); err != nil {
			log.Println(err)
			return nil, err
		}
		userScores = append(userScores, &userScore)
	}
	return userScores, nil
}
