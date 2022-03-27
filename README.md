# Pac

Pac is a opinionated framework based on Fiber to speed up basic monolithic back-end systems.

## Usage

1. Import this package

```
    import "github.com/teampui/pac"
```

2. Then create applications

```
func main() {
	app := pac.NewApp(
		pac.UseListenPortFromEnv(":3000"), // with default to 3000 if not set
		pac.UseLogger(),
	)

	app.Add(NewRootpathService())

	app.Start()
}
```

3. Create Root service

```
package main

import 	"github.com/gofiber/fiber/v2"

type RootpathService struct{}

func NewRootpathService() *RootpathService {
	return new(RootpathService)
}

func (s *RootpathService) Register(app *pac.App, r *fiber.App) {
	r.Get("/", s.HandleRootpath)
}

func (s *RootpathService) HandleRootpath(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

```
