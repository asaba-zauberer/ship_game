package usecase

import (
	"20dojo-online/constant"
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Create(name string) (authToken string, err error)
	SelectByAuthToken(authToken string) (*model.User, error)
	Update(record *model.User, name string, coin, stage int32) error
	GetHighScoresByID(id string) ([]highScoreList, error)
}

type userUsecase struct {
	userRepo      repository.UserRepo
	userScoreRepo repository.UserScoreRepo
}

type highScoreList struct {
	Stage int32
	Score int32
}

// NewUseCase UserUseCaseを生成
func NewUserUsecase(userRepo repository.UserRepo, userScoreRepo repository.UserScoreRepo) UserUsecase {
	return &userUsecase{
		userRepo:      userRepo,
		userScoreRepo: userScoreRepo,
	}
}

// CreateUser Userを新規作成するためのユースケース
func (uc *userUsecase) Create(name string) (string, error) {

	// UUIDでユーザIDを生成
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// UUIDで認証トークンを生成
	authToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	err = uc.userRepo.Create(userID.String(), authToken.String(), name, constant.InitialCoin, constant.InitialStage)
	if err != nil {
		return "", err
	}

	return authToken.String(), nil
}

// SelectByAuthToken 認証トークンからUserを取得するためのユースケース
func (uc *userUsecase) SelectByAuthToken(authToken string) (*model.User, error) {
	user, err := uc.userRepo.SelectByAuthToken(authToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update Userを更新するためのユースケース
func (uc *userUsecase) Update(record *model.User, name string, coin, stage int32) error {
	if err := uc.userRepo.Update(record, name, coin, stage); err != nil {
		return err
	}
	return nil
}

// GetHighScoresByID ユーザーの各ステージごとのハイスコアを取得するためのユースケース
func (uc *userUsecase) GetHighScoresByID(id string) ([]highScoreList, error) {
	userHighScore, err := uc.userScoreRepo.GetScoresByID(id)
	if err != nil {
		return nil, err
	}

	userHighScores := make([]highScoreList, len(userHighScore))
	for i, scoreData := range userHighScore {
		userHighScores[i] = highScoreList{
			Stage: scoreData.Stage,
			Score: scoreData.Score,
		}

	}

	return userHighScores, nil
}
