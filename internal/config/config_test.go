package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Setenv("JWT_SECRET", "904d6807fcd82fbed6e2289a8ff52178")
	cfg, err := LoadConfig()
	require.NotEmpty(t, cfg)
	require.NoError(t, err)

	assert.Equal(t, "904d6807fcd82fbed6e2289a8ff52178", cfg.JWTSecret)
}

func TestLoadConfig_InvalidData(t *testing.T) {
	tests := []struct {
		Name  string
		Key   string
		Value string
	}{
		{Name: "Invalid JWT secret", Key: "JWT_SECRET", Value: "test"},
		{Name: "Invalid port", Key: "PORT", Value: "-1"},
		{Name: "Invalid app env", Key: "APP_ENV", Value: "test"},
		{Name: "Invalid timeout", Key: "TIMEOUT", Value: "-1s"},
		{Name: "Invalid idle timeout", Key: "IDLE_TIMEOUT", Value: "-1s"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv(test.Key, test.Value)
			_, err := LoadConfig()
			require.Error(t, err)
		})
	}
}
