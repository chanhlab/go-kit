package logger

import (
	"context"
	"testing"
)

func TestLogger(_ *testing.T) {
	NewDefault()
	ctx := context.Background()
	ctx.Value("test")
	logger := WithTracerContext(ctx, Log)
	logger.Info("test")
}
