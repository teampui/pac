package pac

import "embed"

func ProvideTemplateFS(fs embed.FS) AppOption {
	return func(a *App) {
		a.templateFs = &fs
	}
}
