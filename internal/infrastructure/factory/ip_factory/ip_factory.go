package ip_factory

import (
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/ip_handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
	"github.com/redis/go-redis/v9"
)

func NewCreateIpRepositoryFactory(redisDatabase *redis.Client) *ip_repository.IPRepository {
	return ip_repository.NewIPRepository(*redisDatabase)
}

func NewCreateIpRuleHandlerFactory(redisDatabase *redis.Client) ip_handlers.CreateIpRuleHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(redisDatabase)
	createIpUseCase := ip_usecase.NewCreateIpRulesUseCase(*ipRepository)
	createIPRuleHandler := ip_handlers.NewCreateIpRuleHandler(createIpUseCase)
	return createIPRuleHandler
}

func NewGetIpRuleByIpHandlerFactory(redisDatabase *redis.Client) ip_handlers.GetIpRuleByIPHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(redisDatabase)
	getIpRuleByIpUsecase := ip_usecase.NewGetIpRulesByIpUseCase(*ipRepository)
	getIpRuleByIpHandler := ip_handlers.NewGetIpRuleByIPHandler(getIpRuleByIpUsecase)
	return getIpRuleByIpHandler
}

func NewGetAllIPRulesHandlerFactory(redisDatabase *redis.Client) ip_handlers.GetAllIPRulesHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(redisDatabase)
	getAllIpRulesUsecase := ip_usecase.NewGetAllIPUseCase(ipRepository)
	getAllIPRulesHandler := ip_handlers.NewGetAllIPRulesHandler(getAllIpRulesUsecase)
	return getAllIPRulesHandler
}

func NewDeleteIPRuleHandlerFactory(redisDatabase *redis.Client) ip_handlers.DeleteIPRuleHandler {
	ipRepository := NewCreateIpRepositoryFactory(redisDatabase)
	deleteUpRuleUseCase := ip_usecase.NewDeleteIPRulesUseCase(*ipRepository)
	deleteIPRuleHandler := ip_handlers.NewDeleteIPRuleHandler(deleteUpRuleUseCase)
	return deleteIPRuleHandler
}

func NewUpdateIPRuleHandlerFactory(redisDatabase *redis.Client) ip_handlers.UpdateIPRuleHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(redisDatabase)
	updateRuleUsecase := ip_usecase.NewUpdateIpRulesByIdUseCase(*ipRepository)
	updateIPRuleHandler := ip_handlers.NewUpdateIPRuleHandler(updateRuleUsecase)
	return updateIPRuleHandler
}
