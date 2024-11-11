package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/config"
	"github.com/masilvasql/go-rate-limiter/internal/adapters/middleware"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/ip_handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/admin/handlers/token_handlers"
	app "github.com/masilvasql/go-rate-limiter/internal/infrastructure/app/handlers"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/middlewares"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
	"github.com/masilvasql/go-rate-limiter/pkg"
)

func main() {
	rootPath := pkg.GetRootPath()
	envConfig, err := config.LoadConfig(rootPath)
	if err != nil {
		panic(errors.New("Error loading config"))
	}

	redisDatabase, err := redis.NewRedisClient(
		envConfig.RedisHost,
		envConfig.RedisPort,
		envConfig.RedisPassword,
		envConfig.RedisDB)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	rateLimiter := middleware.NewRateLimiter(envConfig.MaximumLimitRequestPerSecond, envConfig.LimitedByIP, envConfig.LimitedByToken, envConfig.ExpiresIn)
	client := r.Group("/app")
	{
		client.Use(middlewares.RateLimiter(rateLimiter))
		client.GET("/hello", app.HelloHandler)
		client.GET("/bye", app.ByHandler)
	}

	ipRepository := ip_repository.NewIPRepository(*redisDatabase)

	createIpUseCase := ip_usecase.NewCreateIpRulesUseCase(*ipRepository)
	createIPRuleHandler := ip_handlers.NewCreateIpRuleHandler(createIpUseCase)

	getIpRuleByIpUsecase := ip_usecase.NewGetIpRulesByIpUseCase(*ipRepository)
	getIpRuleByIpHandler := ip_handlers.NewGetIpRuleByIPHandler(getIpRuleByIpUsecase)

	deleteUpRuleUseCase := ip_usecase.NewDeleteIPRulesUseCase(*ipRepository)
	deleteIPRuleHandler := ip_handlers.NewDeleteIPRuleHandler(deleteUpRuleUseCase)

	updateRuleUsecase := ip_usecase.NewUpdateIpRulesByIdUseCase(*ipRepository)
	updateIPRuleHandler := ip_handlers.NewUpdateIPRuleHandler(updateRuleUsecase)

	tokenRepository := token_repository.NewTokenRepository(*redisDatabase)

	createTokenUseCase := token_usecase.NewCreateTokenRulesUseCase(*tokenRepository)
	createTokenRuleHandler := token_handlers.NewCreateTokenRuleHandler(createTokenUseCase)

	getTokenRuleByTokenUsecase := token_usecase.NewGetTokenRulesByTokenUseCase(*tokenRepository)
	getTokenRuleByTokenHandler := token_handlers.NewGetTokenRuleByTokenHandler(getTokenRuleByTokenUsecase)

	deleteTokenRuleUseCase := token_usecase.NewDeleteTokenRulesUseCase(*tokenRepository)
	deleteTokenRuleHandler := token_handlers.NewDeleteTokenRuleHandler(deleteTokenRuleUseCase)

	updateTokenRuleUsecase := token_usecase.NewUpdateTokenRulesByIdUseCase(*tokenRepository)
	updateTokenRuleHandler := token_handlers.NewUpdateTokenRuleHandler(updateTokenRuleUsecase)

	admin := r.Group("/admin")
	{
		ip := admin.Group("/ip")
		{
			ip.POST("/ip-rule", createIPRuleHandler.Handle)
			ip.GET("/ip-rule/:ip", getIpRuleByIpHandler.Handle)
			ip.DELETE("/ip-rule/:id", deleteIPRuleHandler.Handle)
			ip.PUT("/ip-rule/:id", updateIPRuleHandler.Handle)
		}

		token := admin.Group("/token")
		{
			token.POST("/token-rule", createTokenRuleHandler.Handle)
			token.GET("/token-rule/:token", getTokenRuleByTokenHandler.Handle)
			token.DELETE("/token-rule/:id", deleteTokenRuleHandler.Handle)
			token.PUT("/token-rule/:id", updateTokenRuleHandler.Handle)
		}
	}

	fmt.Println("Server running on port: ", envConfig.ServerPort)
	if err := r.Run(":" + envConfig.ServerPort); err != nil {
		panic(err)
	}
}
