package format

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// ConfigFileName is the name of the configuration file.
const ConfigFileName = ".bqfmt.toml"

// Config represents a .bqfmt.toml configuration file.
type Config struct {
	// DefaultStyle is the name of the style to use when no modeline
	// is present.
	DefaultStyle string `toml:"default_style"`
	// LogFile is an optional path to a log file for formatter
	// diagnostics.
	LogFile string `toml:"log_file,omitempty"`
	// Styles is the list of named formatting styles.
	Styles []Style `toml:"styles"`
}

// DefaultConfig returns a Config with the built-in "default" and "raw"
// styles and "default" as the default style.
func DefaultConfig() *Config {
	return &Config{
		DefaultStyle: "default",
		Styles:       DefaultStyles(),
	}
}

// LoadConfig reads and decodes a .bqfmt.toml file at the given path.
// It validates that DefaultStyle references an existing style name.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid config %s: %w", path, err)
	}

	// Initialize all style options.
	for i := range cfg.Styles {
		cfg.Styles[i].Options.Init()
	}

	return &cfg, nil
}

// FindConfig walks from startDir upward through parent directories
// looking for a .bqfmt.toml file.  Returns DefaultConfig() if none
// is found.
func FindConfig(startDir string) (*Config, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return nil, fmt.Errorf("resolving path %s: %w", startDir, err)
	}

	for {
		candidate := filepath.Join(dir, ConfigFileName)

		if _, err := os.Stat(candidate); err == nil {
			return LoadConfig(candidate)
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root without finding a config.
			return DefaultConfig(), nil
		}

		dir = parent
	}
}

// ResolveStyle looks up a style by name.  Returns an error if the
// style is not found.
func (c *Config) ResolveStyle(name string) (*Style, error) {
	for i := range c.Styles {
		if c.Styles[i].Name == name {
			return &c.Styles[i], nil
		}
	}

	available := make([]string, len(c.Styles))
	for i, s := range c.Styles {
		available[i] = s.Name
	}

	return nil, fmt.Errorf("style %q not found (available: %v)", name, available)
}

// DefaultOptions resolves the default style and returns a copy of its
// Options.
func (c *Config) DefaultOptions() (*Options, error) {
	style, err := c.ResolveStyle(c.DefaultStyle)
	if err != nil {
		return nil, fmt.Errorf("resolving default style: %w", err)
	}

	opts := style.Options // copy
	opts.Init()

	return &opts, nil
}

// validate checks internal consistency of a loaded config.
func (c *Config) validate() error {
	if c.DefaultStyle == "" {
		return fmt.Errorf("default_style must not be empty")
	}

	seen := make(map[string]struct{}, len(c.Styles))
	for _, s := range c.Styles {
		if s.Name == "" {
			return fmt.Errorf("every style must have a name")
		}

		if _, ok := seen[s.Name]; ok {
			return fmt.Errorf("duplicate style name %q", s.Name)
		}

		seen[s.Name] = struct{}{}
	}

	if _, ok := seen[c.DefaultStyle]; !ok {
		return fmt.Errorf(
			"default_style %q does not match any defined style",
			c.DefaultStyle,
		)
	}

	return nil
}
