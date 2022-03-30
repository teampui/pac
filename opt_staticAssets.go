package pac

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func ServeStatic(fs embed.FS) AppOption {
	return func(a *App) {
		a.HookCreated = append(a.HookCreated, func(a *App) {
			a.fiber.Use("/public", filesystem.New(filesystem.Config{
				Root: http.FS(fs),
			}))
		})
	}
}
