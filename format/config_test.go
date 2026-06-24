package format_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulourio/gsql/format"
)

func TestDefaultConfig(t *testing.T) {
	cfg := format.DefaultConfig()
	assert.Equal(t, "default", cfg.DefaultStyle)
	assert.Len(t, cfg.Styles, 2)
	assert.Equal(t, "default", cfg.Styles[0].Name)
	assert.Equal(t, "raw", cfg.Styles[1].Name)
}

func TestDefaultConfig_ResolveDefault(t *testing.T) {
	cfg := format.DefaultConfig()
	opts, err := cfg.DefaultOptions()
	require.NoError(t, err)
	assert.Equal(t, format.UpperCase, opts.KeywordStyle)
}

func TestResolveStyle_Found(t *testing.T) {
	cfg := format.DefaultConfig()
	style, err := cfg.ResolveStyle("raw")
	require.NoError(t, err)
	assert.Equal(t, "raw", style.Name)
}

func TestResolveStyle_NotFound(t *testing.T) {
	cfg := format.DefaultConfig()
	_, err := cfg.ResolveStyle("nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonexistent")
}

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "mystyle"

[[styles]]
name = "mystyle"
[styles.options]
keyword_style = "LOWER_CASE"
soft_max_cols = 80
indent_with_entries = false
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	cfg, err := format.LoadConfig(configPath)
	require.NoError(t, err)

	assert.Equal(t, "mystyle", cfg.DefaultStyle)
	require.Len(t, cfg.Styles, 1)
	assert.Equal(t, "mystyle", cfg.Styles[0].Name)
	assert.Equal(t, format.LowerCase, cfg.Styles[0].Options.KeywordStyle)
	assert.Equal(t, 80, cfg.Styles[0].Options.SoftMaxColumns)
	assert.Equal(t, false, cfg.Styles[0].Options.IndentWithEntries)
}

func TestLoadConfig_WithLogFile(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "default"
log_file = "/tmp/bqfmt.log"

[[styles]]
name = "default"
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	cfg, err := format.LoadConfig(configPath)
	require.NoError(t, err)
	assert.Equal(t, "/tmp/bqfmt.log", cfg.LogFile)
}

func TestLoadConfig_DuplicateStyleName(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "dup"

[[styles]]
name = "dup"

[[styles]]
name = "dup"
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	_, err = format.LoadConfig(configPath)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate")
}

func TestLoadConfig_DefaultStyleMissing(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "missing"

[[styles]]
name = "other"
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	_, err = format.LoadConfig(configPath)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing")
}

func TestLoadConfig_EmptyDefaultStyle(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = ""

[[styles]]
name = "something"
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	_, err = format.LoadConfig(configPath)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty")
}

func TestFindConfig_InStartDir(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "s1"

[[styles]]
name = "s1"
[styles.options]
keyword_style = "LOWER_CASE"
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	cfg, err := format.FindConfig(dir)
	require.NoError(t, err)
	assert.Equal(t, "s1", cfg.DefaultStyle)
}

func TestFindConfig_InParentDir(t *testing.T) {
	parent := t.TempDir()
	child := filepath.Join(parent, "subdir")
	err := os.MkdirAll(child, 0o755)
	require.NoError(t, err)

	configPath := filepath.Join(parent, format.ConfigFileName)
	content := `
default_style = "parent_style"

[[styles]]
name = "parent_style"
`
	err = os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	cfg, err := format.FindConfig(child)
	require.NoError(t, err)
	assert.Equal(t, "parent_style", cfg.DefaultStyle)
}

func TestFindConfig_NoConfig(t *testing.T) {
	dir := t.TempDir()

	cfg, err := format.FindConfig(dir)
	require.NoError(t, err)
	// Should return DefaultConfig.
	assert.Equal(t, "default", cfg.DefaultStyle)
	assert.Len(t, cfg.Styles, 2)
}

func TestLoadConfig_MultipleStyles(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, format.ConfigFileName)

	content := `
default_style = "compact"

[[styles]]
name = "compact"
[styles.options]
soft_max_cols = 80
indentation = 4

[[styles]]
name = "wide"
[styles.options]
soft_max_cols = 200
indentation = 2
`
	err := os.WriteFile(configPath, []byte(content), 0o644)
	require.NoError(t, err)

	cfg, err := format.LoadConfig(configPath)
	require.NoError(t, err)

	require.Len(t, cfg.Styles, 2)
	assert.Equal(t, 80, cfg.Styles[0].Options.SoftMaxColumns)
	assert.Equal(t, 4, cfg.Styles[0].Options.Indentation)
	assert.Equal(t, 200, cfg.Styles[1].Options.SoftMaxColumns)
}
