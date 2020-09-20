package injector

import (
	"20dojo-online/pkg/domain/repository"
	"20dojo-online/pkg/infra/mysql"
	repositoryimpl "20dojo-online/pkg/infra/mysql/repository"
	"20dojo-online/pkg/interfaces/handler"
	"20dojo-online/pkg/interfaces/middleware"
	"20dojo-online/pkg/usecase"
)

// -----Middlewareの依存解決-----
func InjectMiddleware() middleware.Middleware {
	return middleware.NewMiddleware(InjectUserUsecase())
}

// -----Userの依存解決-----
func InjectUserRepo() repository.UserRepo {
	return repositoryimpl.NewUserRepo(*mysql.OriginalSqlHandler)
}

func InjectUserScoreRepo() repository.UserScoreRepo {
	return repositoryimpl.NewUserScoreRepo(*mysql.OriginalSqlHandler)
}

func InjectUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(
		InjectUserRepo(),
		InjectUserScoreRepo(),
	)
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}

// -----Gameの依存解決-----
func InjectGameUsecase() usecase.GameUsecase {
	return usecase.NewGameUsecase(InjectUserRepo(), InjectUserScoreRepo())
}

func InjectGameHandler() handler.GameHandler {
	return handler.NewGameHandler(InjectGameUsecase())
}

// -----Rankingの依存解決-----
func InjectRankingUsecase() usecase.RankingUsecase {
	return usecase.NewRankingUsecase(InjectUserRepo(), InjectUserScoreRepo())
}

func InjectRankingHandler() handler.RankingHandler {
	return handler.NewRankingHandler(InjectRankingUsecase())
}

// -----Collectionの依存解決-----
func InjectCollectionItemRepo() repository.CollectionItemRepo {
	return repositoryimpl.NewCollectionItemRepo(*mysql.OriginalSqlHandler)
}

func InjectUserCollectionItemRepo() repository.UserCollectionItemRepo {
	return repositoryimpl.NewUserCollectionItemRepo(*mysql.OriginalSqlHandler)
}

func InjectCollectionUsecase() usecase.CollectionUsecase {
	return usecase.NewCollectionUsecase(InjectCollectionItemRepo(), InjectUserCollectionItemRepo())
}

func InjectCollectionHandler() handler.CollectionHandler {
	return handler.NewCollectionHandler(InjectCollectionUsecase())
}

// -----Gachaの依存解決-----
func InjectGachaProbabilityRepo() repository.GachaProbabilityRepo {
	return repositoryimpl.NewGachaProbabilityRepo(*mysql.OriginalSqlHandler)
}

func InjectGachaUsecase() usecase.GachaUsecase {
	return usecase.NewGachaUsecase(
		InjectUserRepo(),
		InjectGachaProbabilityRepo(),
		InjectCollectionItemRepo(),
		InjectUserCollectionItemRepo(),
	)
}

func InjectGachaHandler() handler.GachaHandler {
	return handler.NewGachaHandler(InjectGachaUsecase())
}
