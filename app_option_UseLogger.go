package pac

import "github.com/gofiber/fiber/v2/middleware/logger"

func UseLogger() AppOption {
	return func(a *App) {
		a.hookCreated = append(a.hookCreated, func(a *App) {
			a.fiber.Use(logger.New())
		})
	}
}
