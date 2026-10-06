package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Autobuska-Stanica/Carrier-Sync-Proxy/internal/config"
	"go.opentelemetry.io/otel/trace"
)

type OtelHandler struct {
	slog.Handler
}

func NewOtelHandler(h slog.Handler) *OtelHandler {
	return &OtelHandler{Handler: h}
}

func (h *OtelHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Automatically attach trace_id and span_id to the log record
		r.AddAttrs(
			slog.String("trace_id", span.SpanContext().TraceID().String()),
			slog.String("span_id", span.SpanContext().SpanID().String()),
		)
	}
	return h.Handler.Handle(ctx, r)
}

func New(cfg *config.Config) *slog.Logger {
	stream := fmt.Sprintf("app=%s,env=%s", cfg.AppName, cfg.Env)

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs([]slog.Attr{
		slog.String("_stream", stream),
	})

	return slog.New(NewOtelHandler(jsonHandler))
}
