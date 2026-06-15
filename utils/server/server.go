package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/likhithkp/ping/utils/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}

func RunHttpServer(lc fx.Lifecycle, env *config.Env, app *fiber.App, logger *zap.Logger) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Success")
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("[server.go] Starting Fiber HTTP server", zap.String("port", env.Port))
				if err := app.Listen(env.Port); err != nil {
					logger.Fatal("[server.go] Failed to start Fiber server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down Fiber HTTP server")
			return app.Shutdown()
		},
	})
}
