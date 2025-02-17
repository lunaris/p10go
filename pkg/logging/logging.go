package logging

import (
	"context"
	"log/slog"
)

type Logger interface {
	Debugf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
}

type slogLogger struct {
	ctx    context.Context
	logger *slog.Logger
}

func NewSlogLogger(ctx context.Context, logger *slog.Logger) Logger {
	return &slogLogger{
		ctx:    ctx,
		logger: logger,
	}
}

func (l *slogLogger) Debugf(message string, args ...interface{}) {
	l.logger.Log(l.ctx, slog.LevelDebug, message, args...)
}

func (l *slogLogger) Infof(message string, args ...interface{}) {
	l.logger.Log(l.ctx, slog.LevelInfo, message, args...)
}

func (l *slogLogger) Warnf(message string, args ...interface{}) {
	l.logger.Log(l.ctx, slog.LevelWarn, message, args...)
}

func (l *slogLogger) Errorf(message string, args ...interface{}) {
	l.logger.Log(l.ctx, slog.LevelError, message, args...)
}
