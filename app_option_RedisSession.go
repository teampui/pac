package pac

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

func UseRedisSession(cfg RedisSessionConfig) AppOption {

	if cfg.ClientKeystore == "" {
		panic("[pac/redisSession] cannot start, missed ClientKeystore settings")
	}

	if cfg.RedisURL == "" {
		panic("[pac/redisSession] cannot start, missed RedisURL settings")
	}

	if cfg.Expiration == 0 {
		cfg.Expiration = 6 * time.Hour
	}

	return func(a *App) {
		a.hookCreated = append(a.hookCreated, func(a *App) {
			a.fiber.Use(SessionInjector(
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

type RedisSessionConfig struct {
	ClientKeystore string
	RedisURL       string
	Expiration     time.Duration
}
