package app

import (
	"os"
	"testing"

	"github.com/TapokGo/tapok-drive/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	t.Setenv("JWT_SECRET", "7test8a0c21f3ebb9testb749test7fb")
	t.Setenv("APP_ENV", "prod")

	tmpFile, err := os.CreateTemp("", "ap-log-*.log")
	require.NoError(t, err)

	t.Setenv("LOG_PATH", tmpFile.Name())

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	require.NotEmpty(t, cfg)

	app, err := New(cfg)
	require.NoError(t, err)
	require.NotEmpty(t, app)

	err = app.Close()
	require.NoError(t, err)
	assert.Equal(t, app.Logger, nil)
}
