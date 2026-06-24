package format

import (
	"fmt"
	"strings"
)

const (
	modelinePrefixSlash = "// bqfmt:"
	modelinePrefixDash  = "-- bqfmt:"
)

// Modeline is the parsed result of a bqfmt modeline comment.
type Modeline struct {
	// Skip is true when the "skip" directive is present, meaning
	// the file should not be formatted.
	Skip bool
	// StyleName is the name of the style to use as a base, from a
	// "style=..." directive.  Empty means use the config default.
	StyleName string
	// Overrides contains raw key=value pairs that patch on top of the
	// resolved style.  Keys are TOML field names (e.g. "keyword_style").
	Overrides map[string]string
}

// ParseModeline parses a single "// bqfmt: ..." or "-- bqfmt: ..." line.
// Returns (nil, nil) if the line is not a modeline.
func ParseModeline(line string) (*Modeline, error) {
	trimmed := strings.TrimSpace(line)

	var body string

	switch {
	case strings.HasPrefix(trimmed, modelinePrefixSlash):
		body = strings.TrimSpace(trimmed[len(modelinePrefixSlash):])
	case strings.HasPrefix(trimmed, modelinePrefixDash):
		body = strings.TrimSpace(trimmed[len(modelinePrefixDash):])
	default:
		return nil, nil
	}

	if body == "" {
		return nil, fmt.Errorf("empty modeline")
	}

	ml := &Modeline{
		Overrides: make(map[string]string),
	}

	directives := strings.Split(body, ",")
	for _, d := range directives {
		d = strings.TrimSpace(d)
		if d == "" {
			continue
		}

		if err := ml.parseDirective(d); err != nil {
			return nil, fmt.Errorf("invalid modeline directive %q: %w", d, err)
		}
	}

	return ml, nil
}

// parseDirective parses a single directive like "skip", "style=raw",
// or "keyword_style=lowercase".
func (ml *Modeline) parseDirective(d string) error {
	if d == "skip" {
		ml.Skip = true
		return nil
	}

	key, value, ok := strings.Cut(d, "=")
	if !ok {
		return fmt.Errorf("expected key=value or 'skip', got %q", d)
	}

	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	if key == "" {
		return fmt.Errorf("empty key in directive")
	}

	if value == "" {
		return fmt.Errorf("empty value for key %q", key)
	}

	if key == "style" {
		ml.StyleName = value
		return nil
	}

	ml.Overrides[key] = value

	return nil
}

// ExtractModeline scans the first non-empty lines of input looking
// for a modeline.  Returns (nil, nil) if no modeline is found.
func ExtractModeline(input string) (*Modeline, error) {
	lines := strings.SplitN(input, "\n", 10) // only look at first few lines

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// Only consider comment lines at the top of the file.
		if !strings.HasPrefix(trimmed, "//") && !strings.HasPrefix(trimmed, "--") {
			// Reached a non-comment, non-empty line → stop scanning.
			return nil, nil
		}

		if strings.HasPrefix(trimmed, modelinePrefixSlash) ||
			strings.HasPrefix(trimmed, modelinePrefixDash) {
			return ParseModeline(trimmed)
		}
	}

	return nil, nil
}

// ApplyModeline resolves a Modeline against a Config to produce final
// Options.
//
//  1. If ml is nil or has no StyleName, use cfg.DefaultStyle.
//  2. Resolve the style from the config.
//  3. Apply any overrides from the modeline.
func ApplyModeline(cfg *Config, ml *Modeline) (*Options, error) {
	styleName := cfg.DefaultStyle
	var overrides map[string]string

	if ml != nil && ml.StyleName != "" {
		styleName = ml.StyleName
	}

	if ml != nil {
		overrides = ml.Overrides
	}

	style, err := cfg.ResolveStyle(styleName)
	if err != nil {
		return nil, fmt.Errorf("applying modeline: %w", err)
	}

	opts := style.Options // copy
	opts.Init()

	if len(overrides) > 0 {
		if err := opts.ApplyOverrides(overrides); err != nil {
			return nil, fmt.Errorf("applying modeline overrides: %w", err)
		}
	}

	return &opts, nil
}
