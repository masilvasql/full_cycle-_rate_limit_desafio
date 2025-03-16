package middleware_test

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/adapters/middleware"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockIpRepository struct {
	mock.Mock
}

func (m *MockIpRepository) Create(ctx context.Context, ipEntity *entity.IPEntity) error {
	args := m.Called(ctx, ipEntity)
	return args.Error(0)
}

func (m *MockIpRepository) GetKey(ctx context.Context, ip string) (*entity.IPEntity, error) {
	args := m.Called(ctx, ip)
	return args.Get(0).(*entity.IPEntity), args.Error(1)
}

func (m *MockIpRepository) GetById(ctx context.Context, id string) (*entity.IPEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.IPEntity), args.Error(1)
}

func (m *MockIpRepository) GetAll(ctx context.Context) ([]entity.IPEntity, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.IPEntity), args.Error(1)
}

func (m *MockIpRepository) Update(ctx context.Context, ipEntity entity.IPEntity) error {
	args := m.Called(ctx, ipEntity)
	return args.Error(0)
}

func (m *MockIpRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockTokenRepository struct {
	mock.Mock
}

func (m *MockTokenRepository) Create(ctx context.Context, tokenEntity *entity.TokenEntity) (*entity.TokenEntity, error) {
	args := m.Called(ctx, tokenEntity)
	return args.Get(0).(*entity.TokenEntity), args.Error(1)
}

func (m *MockTokenRepository) GetByToken(ctx context.Context, token string) (*entity.TokenEntity, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*entity.TokenEntity), args.Error(1)
}

func (m *MockTokenRepository) GetById(ctx context.Context, id string) (*entity.TokenEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.TokenEntity), args.Error(1)
}

func (m *MockTokenRepository) GetAll(ctx context.Context) ([]entity.TokenEntity, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.TokenEntity), args.Error(1)
}

func (m *MockTokenRepository) Update(ctx context.Context, tokenEntity entity.TokenEntity) error {
	args := m.Called(ctx, tokenEntity)
	return args.Error(0)
}

func (m *MockTokenRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockRateLimiterRepository struct {
	mock.Mock
}

func (m *MockRateLimiterRepository) Create(ctx context.Context, requestsKey string, now int64) error {
	args := m.Called(ctx, requestsKey, now)
	return args.Error(0)
}

func (m *MockRateLimiterRepository) FindBanKey(ctx context.Context, tokenIP string) (bool, error) {
	args := m.Called(ctx, tokenIP)
	return args.Bool(0), args.Error(1)
}

func (m *MockRateLimiterRepository) GetTotRequestInPeriod(ctx context.Context, tokenIP string, windowStart int64) (int64, error) {
	args := m.Called(ctx, tokenIP, windowStart)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRateLimiterRepository) AddBanKey(ctx context.Context, tokenIP string, duration string) error {
	args := m.Called(ctx, tokenIP, duration)
	return args.Error(0)
}

func TestCheckRateLimit(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	ip := "192.168.1.1"
	token := "some-token"

	ipEntity := &entity.IPEntity{
		IP:         ip,
		MaxRequest: 10,
		ExpiresIn:  "1h",
	}

	tokenEntity := &entity.TokenEntity{
		Token:      token,
		MaxRequest: 10,
		ExpiresIn:  "1h",
	}

	mockIpRepo.On("GetKey", mock.Anything, ip).Return(ipEntity, nil)
	mockTokenRepo.On("GetByToken", mock.Anything, token).Return(tokenEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, ip).Return(false, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, token).Return(false, nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, ip, mock.Anything).Return(int64(5), nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, token, mock.Anything).Return(int64(5), nil)
	mockRateLimiterRepo.On("Create", mock.Anything, ip, mock.Anything).Return(nil)
	mockRateLimiterRepo.On("Create", mock.Anything, token, mock.Anything).Return(nil)

	err := rateLimiter.CheckRateLimit(ip, "")
	require.NoError(t, err)

	err = rateLimiter.CheckRateLimit("", token)
	require.NoError(t, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)
}

func TestCheckRateLimitBan(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	ip := "192.1.1.1"

	ipEntity := &entity.IPEntity{
		IP:         ip,
		MaxRequest: 10,
	}

	mockIpRepo.On("GetKey", mock.Anything, ip).Return(ipEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, ip).Return(true, nil)

	err := rateLimiter.CheckRateLimit(ip, "")
	require.Error(t, err)
	require.Equal(t, middleware.LimitReached, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)

}

func TestCheckRateLimitBanToken(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	token := "some-token"

	tokenEntity := &entity.TokenEntity{
		Token:      token,
		MaxRequest: 10,
	}

	mockTokenRepo.On("GetByToken", mock.Anything, token).Return(tokenEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, token).Return(true, nil)

	err := rateLimiter.CheckRateLimit("", token)
	require.Error(t, err)
	require.Equal(t, middleware.LimitReached, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)

}

func TestCheckRateLimitNotBanAndNotExceeded(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	ip := "192.1.1.1"

	ipEntity := &entity.IPEntity{
		IP:         ip,
		MaxRequest: 10,
	}

	mockIpRepo.On("GetKey", mock.Anything, ip).Return(ipEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, ip).Return(false, nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, ip, mock.Anything).Return(int64(5), nil)
	mockRateLimiterRepo.On("Create", mock.Anything, ip, mock.Anything).Return(nil)

	err := rateLimiter.CheckRateLimit(ip, "")
	require.NoError(t, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)
}

func TestCheckRateLimitNotBanAndNotExceededToken(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	token := "some-token"

	tokenEntity := &entity.TokenEntity{
		Token:      token,
		MaxRequest: 10,
	}

	mockTokenRepo.On("GetByToken", mock.Anything, token).Return(tokenEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, token).Return(false, nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, token, mock.Anything).Return(int64(5), nil)
	mockRateLimiterRepo.On("Create", mock.Anything, token, mock.Anything).Return(nil)

	err := rateLimiter.CheckRateLimit("", token)
	require.NoError(t, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)
}

func TestCheckIPRateLimitExceededInPeriod(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	ip := "192.1.1.1"

	ipEntity := &entity.IPEntity{
		IP:         ip,
		MaxRequest: 2,
		ExpiresIn:  "1h",
	}

	mockIpRepo.On("GetKey", mock.Anything, ip).Return(ipEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, ip).Return(false, nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, ip, mock.Anything).Return(int64(0), nil).Once()
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, ip, mock.Anything).Return(int64(1), nil).Once()
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, ip, mock.Anything).Return(int64(2), nil).Once()
	mockRateLimiterRepo.On("Create", mock.Anything, ip, mock.Anything).Return(nil).Times(2)
	mockRateLimiterRepo.On("AddBanKey", mock.Anything, ip, ipEntity.ExpiresIn).Return(nil)

	err := rateLimiter.CheckRateLimit(ip, "")
	require.NoError(t, err)

	err = rateLimiter.CheckRateLimit(ip, "")
	require.NoError(t, err)

	err = rateLimiter.CheckRateLimit(ip, "")
	require.Error(t, err)
	require.Equal(t, middleware.LimitReached, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)
}

func TestCheckTokenRateLimitExceededInPeriod(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, true)

	token := "some-token"

	tokenEntity := &entity.TokenEntity{
		Token:      token,
		MaxRequest: 2,
		ExpiresIn:  "1h",
	}

	mockTokenRepo.On("GetByToken", mock.Anything, token).Return(tokenEntity, nil)
	mockRateLimiterRepo.On("FindBanKey", mock.Anything, token).Return(false, nil)
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, token, mock.Anything).Return(int64(0), nil).Once()
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, token, mock.Anything).Return(int64(1), nil).Once()
	mockRateLimiterRepo.On("GetTotRequestInPeriod", mock.Anything, token, mock.Anything).Return(int64(2), nil).Once()
	mockRateLimiterRepo.On("Create", mock.Anything, token, mock.Anything).Return(nil).Times(2)
	mockRateLimiterRepo.On("AddBanKey", mock.Anything, token, tokenEntity.ExpiresIn).Return(nil)

	err := rateLimiter.CheckRateLimit("", token)
	require.NoError(t, err)

	err = rateLimiter.CheckRateLimit("", token)
	require.NoError(t, err)

	err = rateLimiter.CheckRateLimit("", token)
	require.Error(t, err)
	require.Equal(t, middleware.LimitReached, err)

	mockIpRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockRateLimiterRepo.AssertExpectations(t)
}

func TestCheckRateLimitDisableIPRateLimit(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, false, true)

	ip := "123.1.1.1"

	err := rateLimiter.CheckRateLimit(ip, "")
	require.Error(t, middleware.ErrNotAuthorized, err)

}

func TestCheckRateLimitDisableTokenRateLimit(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, true, false)

	token := "some-token"

	err := rateLimiter.CheckRateLimit("", token)
	require.Error(t, middleware.ErrNotAuthorized, err)
}

func TestCheckRateLimitDisableIPAndTokenRateLimit(t *testing.T) {
	mockIpRepo := new(MockIpRepository)
	mockTokenRepo := new(MockTokenRepository)
	mockRateLimiterRepo := new(MockRateLimiterRepository)

	rateLimiter := middleware.NewRateLimiter(mockTokenRepo, mockIpRepo, mockRateLimiterRepo, false, false)

	ip := "1.1.1.1"

	err := rateLimiter.CheckRateLimit(ip, "")
	require.Error(t, middleware.ErrNotAuthorized, err)

	token := "some-token"

	err = rateLimiter.CheckRateLimit("", token)
	require.Error(t, middleware.ErrNotAuthorized, err)

}
