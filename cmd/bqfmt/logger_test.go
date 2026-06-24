package main

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestUnquotedTextHandler(t *testing.T) {
	var buf bytes.Buffer
	handler := NewUnquotedTextHandler(&buf, slog.LevelDebug)
	logger := slog.New(handler)

	// Log with message and multiple key-value pairs
	logger.Info("Hello, World!", "key1", "val1", "key2", "val with spaces")

	output := buf.String()
	t.Logf("Output: %q", output)

	// Verify time format (starts with time=)
	if !strings.Contains(output, "time=") {
		t.Errorf("Expected output to contain time=, got %q", output)
	}

	// Verify level
	if !strings.Contains(output, "level=INFO") {
		t.Errorf("Expected output to contain level=INFO, got %q", output)
	}

	// Verify msg (unquoted!)
	if !strings.Contains(output, "msg=Hello, World!") {
		t.Errorf("Expected output to contain msg=Hello, World!, got %q", output)
	}

	// Verify key1 (unquoted)
	if !strings.Contains(output, "key1=val1") {
		t.Errorf("Expected output to contain key1=val1, got %q", output)
	}

	// Verify key2 (unquoted even with spaces!)
	if !strings.Contains(output, "key2=val with spaces") {
		t.Errorf("Expected output to contain key2=val with spaces, got %q", output)
	}
}

func TestUnquotedTextHandlerWithAttrs(t *testing.T) {
	var buf bytes.Buffer
	handler := NewUnquotedTextHandler(&buf, slog.LevelDebug).WithAttrs([]slog.Attr{
		slog.String("common", "some value"),
	})
	logger := slog.New(handler)

	logger.Warn("A warning message", "specific", "details here")

	output := buf.String()
	t.Logf("Output: %q", output)

	if !strings.Contains(output, "msg=A warning message") {
		t.Errorf("Expected msg=A warning message, got %q", output)
	}
	if !strings.Contains(output, "common=some value") {
		t.Errorf("Expected common=some value, got %q", output)
	}
	if !strings.Contains(output, "specific=details here") {
		t.Errorf("Expected specific=details here, got %q", output)
	}
}
