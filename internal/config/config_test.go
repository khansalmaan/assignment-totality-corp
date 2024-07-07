package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	// Test with environment variable set
	os.Setenv("TEST_KEY", "test_value")
	value := getEnv("TEST_KEY", "default_value")
	assert.Equal(t, "test_value", value)

	// Test with environment variable not set
	os.Unsetenv("TEST_KEY")
	value = getEnv("TEST_KEY", "default_value")
	assert.Equal(t, "default_value", value)
}

func TestLoad(t *testing.T) {
	// Test with environment variable set
	os.Setenv("SERVER_ADDRESS", ":6000")
	cfg := Load()
	assert.Equal(t, ":6000", cfg.Server.Address)

	// Test with environment variable not set
	os.Unsetenv("SERVER_ADDRESS")
	cfg = Load()
	assert.Equal(t, ":50051", cfg.Server.Address)
}
