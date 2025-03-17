package token_factory

import (
	"database/sql"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/token_handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
	"github.com/redis/go-redis/v9"
)

func NewCreateTokenRepositoryFactory(driver string, redisDatabase *redis.Client, db *sql.DB) *token_repository.TokenRepository {
	switch driver {
	case "redis":
		return token_repository.NewTokenRepository(*redisDatabase)
	case "mysql":
		panic("mysql is not implemented yet")
	default:
		panic("driver not found")
	}
}

func NewcreateTokenRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) token_handlers.CreateTokenRuleHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(driver, redisDatabase, db)
	createTokenUseCase := token_usecase.NewCreateTokenRulesUseCase(*tokenRepository)
	createTokenRuleHandler := token_handlers.NewCreateTokenRuleHandler(createTokenUseCase)
	return createTokenRuleHandler
}

func NewGetTokenRuleByTokenHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) token_handlers.GetTokenRuleByTokenHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(driver, redisDatabase, db)
	getTokenRuleByTokenUsecase := token_usecase.NewGetTokenRulesByTokenUseCase(*tokenRepository)
	getTokenRuleByTokenHandler := token_handlers.NewGetTokenRuleByTokenHandler(getTokenRuleByTokenUsecase)
	return getTokenRuleByTokenHandler
}

func NewGetAllTokenHandlersFactory(driver string, redisDatabase *redis.Client, db *sql.DB) token_handlers.GetAllTokenHandlersInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(driver, redisDatabase, db)
	getAllTokenRulesUsecase := token_usecase.NewGetAllTokenRulesUseCase(*tokenRepository)
	getAllTokenRulesHandler := token_handlers.NewGetAllTokenHandlers(getAllTokenRulesUsecase)
	return getAllTokenRulesHandler
}

func NewDeleteTokenRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) token_handlers.DeleteTokenRuleHandler {
	tokenRepository := NewCreateTokenRepositoryFactory(driver, redisDatabase, db)
	deleteTokenRuleUseCase := token_usecase.NewDeleteTokenRulesUseCase(*tokenRepository)
	deleteTokenRuleHandler := token_handlers.NewDeleteTokenRuleHandler(deleteTokenRuleUseCase)
	return deleteTokenRuleHandler
}

func NewUpdateTokenRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) token_handlers.UpdateTokenRuleHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(driver, redisDatabase, db)
	updateTokenRuleUsecase := token_usecase.NewUpdateTokenRulesByIdUseCase(*tokenRepository)
	updateTokenRuleHandler := token_handlers.NewUpdateTokenRuleHandler(updateTokenRuleUsecase)
	return updateTokenRuleHandler
}
