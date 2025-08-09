package app

import (
	"context"
	"fmt"
	"gateway/internal/config"
	"gateway/internal/infrastructure/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	l := logger.New(defaultLevel, os.Stdout)
	ctx = logger.WithLogger(ctx, l)

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Errorw("failed to load config", "error", err)
		return fmt.Errorf("failed to load config: %w", err)
	}
	fmt.Println(cfg)
	l.Info("App starting")

	<-ctx.Done()

	return nil
}
