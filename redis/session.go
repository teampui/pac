package redis

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/teampui/pac"
)

func ProvideSession(cfg SessionConfig) pac.AppOption {

	if cfg.ClientKeystore == "" {
		panic("pac/redis: cannot start, missed ClientKeystore settings")
	}

	if cfg.RedisURL == "" {
		panic("pac/redis: cannot start, missed RedisURL settings")
	}

	if cfg.Expiration == 0 {
		cfg.Expiration = 6 * time.Hour
	}

	return func(a *pac.App) {
		a.HookCreated = append(a.HookCreated, func(a *pac.App) {
			a.Router().Use(pac.SessionInjector(
				session.New(session.Config{
					// cookie related
					CookieHTTPOnly: true,
					CookieSameSite: "strict",
					// session key
					KeyLookup:  cfg.ClientKeystore,
					Expiration: cfg.Expiration,
					// stroage
					Storage: redis.New(redis.Config{
						URL:   cfg.RedisURL,
						Reset: false,
					}),
				}),
			))
		})
	}
}

type SessionConfig struct {
	ClientKeystore string
	RedisURL       string
	Expiration     time.Duration
}
