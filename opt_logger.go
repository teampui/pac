package pac

import "github.com/gofiber/fiber/v2/middleware/logger"

func UseLogger() AppOption {
	return func(a *App) {
		a.HookCreated = append(a.HookCreated, func(a *App) {
			a.fiber.Use(logger.New())
		})
	}
}
