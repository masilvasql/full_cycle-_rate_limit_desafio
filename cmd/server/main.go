package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/config"
	"github.com/masilvasql/go-rate-limiter/internal/adapters/middleware"
	app "github.com/masilvasql/go-rate-limiter/internal/infrastructure/app/handlers"
	msClient "github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/factory/ip_factory"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/factory/rate_limiter_factory"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/factory/token_factory"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/middlewares"
	"github.com/masilvasql/go-rate-limiter/pkg"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func main() {
	rootPath := pkg.GetRootPath()
	envConfig, err := config.LoadConfig(rootPath)
	if err != nil {
		panic(errors.New("Error loading config"))
	}

	redisDatabase := getRedisDatabase(envConfig)

	r := gin.New()
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = f
	gin.DefaultErrorWriter = f
	r.Use(gin.Recovery())

	rateLimiter := middleware.NewRateLimiter(
		token_factory.NewCreateTokenRepositoryFactory(envConfig.Driver, redisDatabase, nil),
		ip_factory.NewCreateIpRepositoryFactory(envConfig.Driver, redisDatabase, nil),
		rate_limiter_factory.NewCreateReateLimiterRepositoryFactory(envConfig.Driver, redisDatabase, nil),
		envConfig.LimitedByIP,
		envConfig.LimitedByToken)

	client := r.Group("/app")
	{
		client.Use(middlewares.RateLimiter(rateLimiter))
		client.GET("/hello", app.HelloHandler)
		client.GET("/bye", app.ByHandler)
	}

	admin := r.Group("/admin")
	{
		ip := admin.Group("/ip")
		{
			ip.POST("/ip-rule", ip_factory.NewCreateIpRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			ip.GET("/ip-rule/:ip", ip_factory.NewGetIpRuleByIpHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			ip.GET("/ip-rule/all", ip_factory.NewGetAllIPRulesHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			ip.DELETE("/ip-rule/:id", ip_factory.NewDeleteIPRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			ip.PUT("/ip-rule/:id", ip_factory.NewUpdateIPRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
		}

		token := admin.Group("/token")
		{
			token.POST("/token-rule", token_factory.NewcreateTokenRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			token.GET("/token-rule/:token", token_factory.NewGetTokenRuleByTokenHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			token.GET("/token-rule/all", token_factory.NewGetAllTokenHandlersFactory(envConfig.Driver, redisDatabase, nil).Handle)
			token.DELETE("/token-rule/:id", token_factory.NewDeleteTokenRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
			token.PUT("/token-rule/:id", token_factory.NewUpdateTokenRuleHandlerFactory(envConfig.Driver, redisDatabase, nil).Handle)
		}
	}

	fmt.Println("Server running on port: ", envConfig.ServerPort)
	if err := r.Run(":" + envConfig.ServerPort); err != nil {
		panic(err)
	}
}

func getRedisDatabase(envConfig *config.Config) *redis.Client {
	redisDatabase, err := msClient.NewRedisClient(
		envConfig.RedisHost,
		envConfig.RedisPort,
		envConfig.RedisPassword,
		envConfig.RedisDB)

	if err != nil {
		panic(err)
	}

	return redisDatabase
}
