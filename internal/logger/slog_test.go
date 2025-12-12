package logger

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/TapokGo/tapok-drive/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	logger, getContent, clenup := newProdLogger(t)
	logger.Error("logger")

	newLogger := logger.With("test", "test")
	require.NotNil(t, newLogger)
	require.NotEqual(t, logger, newLogger)

	newLogger.Error("new logger")

	logFileContent := getContent()

	clenup()

	logs := strings.Split(strings.TrimSpace(logFileContent), "\n")

	assert.NotContains(t, logs[0], `"test":"test"`)
	assert.Contains(t, logs[1], `"test":"test"`)
}

func newProdLogger(t *testing.T) (logger Logger, readFile func() string, clenup func()) {
	t.Setenv("APP_ENV", "prod")
	t.Setenv("JWT_SECRET", "92test0000afcdc161dd26atestb905d")

	tmpFile, err := os.CreateTemp("", "test-app-*.log")
	require.NoError(t, err)

	t.Setenv("LOG_PATH", tmpFile.Name())

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	require.NotEmpty(t, cfg)

	logger, err = NewSlog(cfg.LogPath, cfg.AppEnv)
	require.NoError(t, err)
	require.NotNil(t, logger)

	return logger,
		func() string {
			content, err := os.ReadFile(tmpFile.Name())
			require.NoError(t, err)
			require.NotEmpty(t, content)

			return string(content)
		},
		func() {
			err := logger.(io.Closer).Close()
			require.NoError(t, err)
		}
}
