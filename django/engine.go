package django

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/teampui/pac"
)

func UseDjangoTemplateEngine(fs embed.FS) pac.AppOption {
	return func(a *pac.App) {
		a.HookBeforeCreate = append(a.HookBeforeCreate, func(c *fiber.Config) {
			c.Views = django.NewFileSystem(http.FS(fs), ".html.twig")
		})
	}
}
