package internal

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigPath(t *testing.T) {
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := filepath.Join("/tmp/home", ".kawaiclaw", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithKAWAICLAW_HOME(t *testing.T) {
	t.Setenv("KAWAICLAW_HOME", "/custom/kawaiclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := filepath.Join("/custom/kawaiclaw", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithKAWAICLAW_CONFIG(t *testing.T) {
	t.Setenv("KAWAICLAW_CONFIG", "/custom/config.json")
	t.Setenv("KAWAICLAW_HOME", "/custom/kawaiclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := "/custom/config.json"

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithPICOCLAW_HOME(t *testing.T) {
	t.Setenv("PICOCLAW_HOME", "/custom/picoclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := filepath.Join("/custom/picoclaw", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithPICOCLAW_CONFIG(t *testing.T) {
	t.Setenv("PICOCLAW_CONFIG", "/custom/config.json")
	t.Setenv("PICOCLAW_HOME", "/custom/picoclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := "/custom/config.json"

	assert.Equal(t, want, got)
}

func TestGetConfigPath_Windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-specific HOME behavior varies; run on windows")
	}

	testUserProfilePath := `C:\Users\Test`
	t.Setenv("USERPROFILE", testUserProfilePath)

	got := GetConfigPath()
	want := filepath.Join(testUserProfilePath, ".kawaiclaw", "config.json")

	require.True(t, strings.EqualFold(got, want), "GetConfigPath() = %q, want %q", got, want)
}
