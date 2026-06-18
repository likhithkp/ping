package server

import (
	"context"

	"github.com/gofiber/contrib/websocket"
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

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("[server.go] Starting server", zap.String("port", env.Port))
				if err := app.Listen(env.Port); err != nil {
					logger.Fatal("[server.go] Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down server")
			return app.Shutdown()
		},
	})
}
