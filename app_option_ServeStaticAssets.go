package pac

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func ServeStaticAssets(fs embed.FS) AppOption {
	return func(a *App) {
		a.hookCreated = append(a.hookCreated, func(a *App) {
			a.fiber.Use("/public", filesystem.New(filesystem.Config{
				Root: http.FS(fs),
			}))
		})
	}
}
