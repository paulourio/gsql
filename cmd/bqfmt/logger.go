package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"
)

type UnquotedTextHandler struct {
	w     io.Writer
	level slog.Level
	attrs []slog.Attr
}

func NewUnquotedTextHandler(w io.Writer, level slog.Level) *UnquotedTextHandler {
	return &UnquotedTextHandler{
		w:     w,
		level: level,
	}
}

func (h *UnquotedTextHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *UnquotedTextHandler) Handle(_ context.Context, r slog.Record) error {
	var buf strings.Builder

	// Format time
	if !r.Time.IsZero() {
		buf.WriteString("time=")
		buf.WriteString(r.Time.Format(time.RFC3339Nano))
		buf.WriteByte(' ')
	}

	// Format level
	buf.WriteString("level=")
	buf.WriteString(r.Level.String())
	buf.WriteByte(' ')

	// Format message
	buf.WriteString("msg=")
	buf.WriteString(r.Message)

	// Format pre-configured attrs
	for _, attr := range h.attrs {
		buf.WriteByte(' ')
		buf.WriteString(attr.Key)
		buf.WriteByte('=')
		buf.WriteString(formatValue(attr.Value))
	}

	// Format record attrs
	r.Attrs(func(attr slog.Attr) bool {
		buf.WriteByte(' ')
		buf.WriteString(attr.Key)
		buf.WriteByte('=')
		buf.WriteString(formatValue(attr.Value))
		return true
	})

	buf.WriteByte('\n')

	_, err := h.w.Write([]byte(buf.String()))
	if err != nil {
		return fmt.Errorf("unquoted text handler: %w", err)
	}
	return nil
}

func (h *UnquotedTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	newAttrs = append(newAttrs, h.attrs...)
	newAttrs = append(newAttrs, attrs...)
	return &UnquotedTextHandler{
		w:     h.w,
		level: h.level,
		attrs: newAttrs,
	}
}

func (h *UnquotedTextHandler) WithGroup(_ string) slog.Handler {
	return h
}

func formatValue(v slog.Value) string {
	v = v.Resolve()
	switch v.Kind() { //nolint:exhaustive
	case slog.KindString:
		return v.String()
	case slog.KindTime:
		return v.Time().Format(time.RFC3339Nano)
	case slog.KindGroup:
		attrs := v.Group()
		var parts []string
		for _, a := range attrs {
			parts = append(parts, fmt.Sprintf("%s=%s", a.Key, formatValue(a.Value)))
		}
		return strings.Join(parts, " ")
	default:
		return fmt.Sprintf("%v", v.Any())
	}
}
