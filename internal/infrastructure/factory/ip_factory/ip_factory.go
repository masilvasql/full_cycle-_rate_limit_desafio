package ip_factory

import (
	"database/sql"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/ip_handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
	"github.com/redis/go-redis/v9"
)

func NewCreateIpRepositoryFactory(driver string, redisDatabase *redis.Client, db *sql.DB) *ip_repository.IPRepository {
	switch driver {
	case "redis":
		return ip_repository.NewIPRepository(*redisDatabase)
	case "mysql":
		panic("mysql is not implemented yet")
	default:
		panic("driver not found")
	}
}

func NewCreateIpRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) ip_handlers.CreateIpRuleHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(driver, redisDatabase, db)
	createIpUseCase := ip_usecase.NewCreateIpRulesUseCase(*ipRepository)
	createIPRuleHandler := ip_handlers.NewCreateIpRuleHandler(createIpUseCase)
	return createIPRuleHandler
}

func NewGetIpRuleByIpHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) ip_handlers.GetIpRuleByIPHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(driver, redisDatabase, db)
	getIpRuleByIpUsecase := ip_usecase.NewGetIpRulesByIpUseCase(*ipRepository)
	getIpRuleByIpHandler := ip_handlers.NewGetIpRuleByIPHandler(getIpRuleByIpUsecase)
	return getIpRuleByIpHandler
}

func NewGetAllIPRulesHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) ip_handlers.GetAllIPRulesHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(driver, redisDatabase, db)
	getAllIpRulesUsecase := ip_usecase.NewGetAllIPUseCase(ipRepository)
	getAllIPRulesHandler := ip_handlers.NewGetAllIPRulesHandler(getAllIpRulesUsecase)
	return getAllIPRulesHandler
}

func NewDeleteIPRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) ip_handlers.DeleteIPRuleHandler {
	ipRepository := NewCreateIpRepositoryFactory(driver, redisDatabase, db)
	deleteUpRuleUseCase := ip_usecase.NewDeleteIPRulesUseCase(*ipRepository)
	deleteIPRuleHandler := ip_handlers.NewDeleteIPRuleHandler(deleteUpRuleUseCase)
	return deleteIPRuleHandler
}

func NewUpdateIPRuleHandlerFactory(driver string, redisDatabase *redis.Client, db *sql.DB) ip_handlers.UpdateIPRuleHandlerInterface {
	ipRepository := NewCreateIpRepositoryFactory(driver, redisDatabase, db)
	updateRuleUsecase := ip_usecase.NewUpdateIpRulesByIdUseCase(*ipRepository)
	updateIPRuleHandler := ip_handlers.NewUpdateIPRuleHandler(updateRuleUsecase)
	return updateIPRuleHandler
}
