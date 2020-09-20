package usecase

import (
	"20dojo-online/constant"
	"20dojo-online/pkg/domain/model"
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type GachaUsecase interface {
	DrawGacha(user *model.User, times int32) ([]*gachaResultList, error)
}

type gachaUsecase struct {
	userRepo               repository.UserRepo
	gachaProbabilityRepo   repository.GachaProbabilityRepo
	collectionItemRepo     repository.CollectionItemRepo
	userCollectionItemRepo repository.UserCollectionItemRepo
}

type gachaResultList struct {
	ID     string
	Name   string
	Rarity int32
	IsNew  bool
}

// NewGachaUsecase GachaUsecaseを作成
func NewGachaUsecase(ur repository.UserRepo, gr repository.GachaProbabilityRepo, cr repository.CollectionItemRepo, ucr repository.UserCollectionItemRepo) GachaUsecase {
	return &gachaUsecase{
		userRepo:               ur,
		gachaProbabilityRepo:   gr,
		collectionItemRepo:     cr,
		userCollectionItemRepo: ucr,
	}
}

// DrawGacha
func (gu gachaUsecase) DrawGacha(user *model.User, times int32) ([]*gachaResultList, error) {
	gachaResults := make([]*gachaResultList, times)

	// ガチャ抽選に関するトランザクション
	err := mysql.Transaction(mysql.OriginalSqlHandler.Conn, func(tx *sql.Tx) error {
		// 他のトランザクションからのユー	ザー情報更新のロック
		if err := gu.userRepo.Lock(user); err != nil {
			return err
		}

		// コイン不足確認
		if user.Coin < times*constant.GachaCoinConsumption {
			return errors.New(fmt.Sprintf("lack of coins, coin: %d", user.Coin))
		}

		// ガチャの排出確立の取得
		gachaProb, err := gu.gachaProbabilityRepo.SelectAll()
		if err != nil {
			log.Println(err)
			return err
		}

		// 抽選処理
		gachaResult, err := lottery(gachaProb, times)
		if err != nil {
			return err
		}

		// 入手アイテムの情報を取得
		gotItemsData, err := gu.collectionItemRepo.SelectByID(gachaResult)
		if err != nil {
			return err
		}

		// ユーザーのアイテム所持情報を取得
		userCollectionItems, err := gu.userCollectionItemRepo.SelectByUserID(user.ID)
		if err != nil {
			return err
		}

		// mapに変換
		userCollectionItemsMap := make(map[string]string)
		for _, data := range userCollectionItems {
			userCollectionItemsMap[data.CollectionItemID] = data.UserID
		}

		// DrawGachaの返却値gachaResultsと保存するガチャ結果userGotItemsを作成
		userGotItems := make(model.UserCollectionItems, times)
		for i, gotItemData := range gotItemsData {
			_, ok := userCollectionItemsMap[gotItemData.ID]

			gachaResults[i] = &gachaResultList{
				ID:    gotItemData.ID,
				Name:  gotItemData.Name,
				IsNew: !ok,
			}
			userGotItems[i] = &model.UserCollectionItem{
				UserID:           user.ID,
				CollectionItemID: gotItemData.ID,
			}
		}

		// 入手アイテムの保存
		if err := gu.userCollectionItemRepo.InsertItems(userGotItems); err != nil {
			return err
		}

		// コインの減算処理
		if err := gu.userRepo.Update(user, user.Name, user.Coin-times*constant.GachaCoinConsumption, user.Stage); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return gachaResults, nil
}

// lottery ガチャ抽選処理
func lottery(gachaProb model.GachaProbabilities, times int32) ([]string, error) {
	var gotItems []string
	totalProb := gachaProb[len(gachaProb)-1].Ratio

	for i := int32(0); i < times; i++ {
		target := rand.Int31n(totalProb)
		gotItem, err := itemBinSearch(gachaProb, target)
		if err != nil {
			return nil, err
		}
		gotItems = append(gotItems, gotItem)
	}

	return gotItems, nil
}

// itemBinSearch ガチャ当選アイテムの二分探索
func itemBinSearch(gachaProbs model.GachaProbabilities, target int32) (string, error) {

	start := 0
	end := len(gachaProbs) - 1

	var mid int
	for start < end {
		if end < start {
			return "", errors.New("bin search error")
		}

		mid = (start + end) / 2

		if gachaProbs[mid].Ratio < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return gachaProbs[start].CollectionItemID, nil
}
