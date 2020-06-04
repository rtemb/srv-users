package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	err := os.Setenv("TOKEN_KEY", "123456")
	require.NoError(t, err)
	err = os.Setenv("LOG_LEVEL", "error")
	require.NoError(t, err)

	err = os.Setenv("REDIS_PORT", "1234")
	require.NoError(t, err)
	err = os.Setenv("REDIS_PASSWORD", "password")
	require.NoError(t, err)
	err = os.Setenv("REDIS_MAX_IDLE", "10")
	require.NoError(t, err)
	err = os.Setenv("REDIS_IDLE_TIMEOUT", "30s")
	require.NoError(t, err)

	err = os.Setenv("APP_PORT", "777")
	require.NoError(t, err)
	err = os.Setenv("GRACEFUL_SHUTDOWN_TIMEOUT", "30s")
	require.NoError(t, err)
	err = os.Setenv("WRITE_TIMEOUT", "30s")
	require.NoError(t, err)
	err = os.Setenv("READ_TIMEOUT", "30s")
	require.NoError(t, err)
	err = os.Setenv("IDLE_TIMEOUT", "30s")
	require.NoError(t, err)

	testDuration, err := time.ParseDuration("30s")
	require.NoError(t, err)

	cfg, err := Load()
	require.NoError(t, err)

	assert.Equal(t, "123456", cfg.AppConfig.TokenKey)
	assert.Equal(t, "error", cfg.AppConfig.LogLevel)
	assert.Equal(t, "localhost", cfg.Redis.Host)
	assert.Equal(t, "1234", cfg.Redis.Port)
	assert.Equal(t, "password", cfg.Redis.Password)
	assert.Equal(t, 10, cfg.Redis.MaxIdle)
	assert.Equal(t, testDuration, cfg.Redis.IdleTimeout)
	assert.Equal(t, "8081", cfg.Server.GatewayPort)
	assert.Equal(t, testDuration, cfg.Server.GracefulShutdownTimeout)
	assert.Equal(t, testDuration, cfg.Server.WriteTimeout)
	assert.Equal(t, testDuration, cfg.Server.ReadTimeout)
	assert.Equal(t, testDuration, cfg.Server.IdleTimeout)
}

func TestLoad_Fail(t *testing.T) {
	os.Clearenv()

	cfg, err := Load()
	assert.Nil(t, cfg)
	assert.Error(t, err)
}
