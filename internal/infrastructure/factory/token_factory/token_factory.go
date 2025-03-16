package token_factory

import (
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/token_handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
	"github.com/redis/go-redis/v9"
)

func NewCreateTokenRepositoryFactory(redisDatabase *redis.Client) *token_repository.TokenRepository {
	return token_repository.NewTokenRepository(*redisDatabase)
}

func NewcreateTokenRuleHandlerFactory(redisDatabase *redis.Client) token_handlers.CreateTokenRuleHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(redisDatabase)
	createTokenUseCase := token_usecase.NewCreateTokenRulesUseCase(*tokenRepository)
	createTokenRuleHandler := token_handlers.NewCreateTokenRuleHandler(createTokenUseCase)
	return createTokenRuleHandler
}

func NewGetTokenRuleByTokenHandlerFactory(redisDatabase *redis.Client) token_handlers.GetTokenRuleByTokenHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(redisDatabase)
	getTokenRuleByTokenUsecase := token_usecase.NewGetTokenRulesByTokenUseCase(*tokenRepository)
	getTokenRuleByTokenHandler := token_handlers.NewGetTokenRuleByTokenHandler(getTokenRuleByTokenUsecase)
	return getTokenRuleByTokenHandler
}

func NewGetAllTokenHandlersFactory(redisDatabase *redis.Client) token_handlers.GetAllTokenHandlersInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(redisDatabase)
	getAllTokenRulesUsecase := token_usecase.NewGetAllTokenRulesUseCase(*tokenRepository)
	getAllTokenRulesHandler := token_handlers.NewGetAllTokenHandlers(getAllTokenRulesUsecase)
	return getAllTokenRulesHandler
}

func NewDeleteTokenRuleHandlerFactory(redisDatabase *redis.Client) token_handlers.DeleteTokenRuleHandler {
	tokenRepository := NewCreateTokenRepositoryFactory(redisDatabase)
	deleteTokenRuleUseCase := token_usecase.NewDeleteTokenRulesUseCase(*tokenRepository)
	deleteTokenRuleHandler := token_handlers.NewDeleteTokenRuleHandler(deleteTokenRuleUseCase)
	return deleteTokenRuleHandler
}

func NewUpdateTokenRuleHandlerFactory(redisDatabase *redis.Client) token_handlers.UpdateTokenRuleHandlerInterface {
	tokenRepository := NewCreateTokenRepositoryFactory(redisDatabase)
	updateTokenRuleUsecase := token_usecase.NewUpdateTokenRulesByIdUseCase(*tokenRepository)
	updateTokenRuleHandler := token_handlers.NewUpdateTokenRuleHandler(updateTokenRuleUsecase)
	return updateTokenRuleHandler
}
