package pac

import (
	"os"
	"strings"
)

func UseListenPortFromEnv(default_listen string) AppOption {
	return func(a *App) {
		envPort := os.Getenv("PORT")
		envPort = strings.TrimSpace(envPort)

		a.port = envPort

		if a.port == "" {
			a.port = default_listen
		}
	}
}
