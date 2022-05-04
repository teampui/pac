package pac

import "github.com/gofiber/fiber/v2/middleware/compress"

func UseCompress() AppOption {
	return func(a *App) {
		a.HookCreated = append(a.HookCreated, func(a *App) {
			a.fiber.Use(compress.New())
		})
	}
}
