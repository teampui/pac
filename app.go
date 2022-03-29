package pac

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

// NewApplication returns a Pac App
func NewApp(opts ...AppOption) *App {
	// create new pac app
	pacApp := App{
		hookCreated: []AppOption{},
		services:    make(map[string]Middleware),
	}

	// apply options
	for _, opt := range opts {
		opt(&pacApp)
	}

	// then we init app
	fiberCfg := fiber.Config{
		DisableStartupMessage: true,
	}

	if pacApp.templateFs != nil {
		fiberCfg.Views = django.NewFileSystem(http.FS(*pacApp.templateFs), ".html.twig")
	}

	// Pass the engine to the Views
	pacApp.fiber = fiber.New(fiberCfg)

	// apply hook
	for _, hook := range pacApp.hookCreated {
		hook(&pacApp)
	}

	// return to caller
	return &pacApp
}

// App is Pac App which accept services and serves
type App struct {
	//
	fiber      *fiber.App
	templateFs *embed.FS
	//
	port string
	//
	hookCreated []AppOption
	services    map[string]Middleware
}

// Add will add service to pac app, callin' its `Register()` to inject routes
// and register service functions to app.
func (a *App) Add(svc Service) {
	svc.Register(a, a.fiber)
}

// Start will start server and listen to given PORT
func (a *App) Start() {
	// check if there is no PORT, then panic errors
	if a.port == "" {
		panic("[pac] cannot start, missed env PORT")
	}

	// if everything ready, start listen request
	err := a.fiber.Listen(a.port)

	if err != nil {
		panic("[pac] cannot start, because error: " + err.Error())
	}
}

// Service return function by its register name, return nil if not found
func (a *App) Middleware(svcName string) Middleware {
	svc, ok := a.services[svcName]

	if !ok {
		return nil
	}

	return svc
}

// RegisterService register function to pac app, get by call `Service()`
func (a *App) RegisterMiddleware(svcName string, svcFunc Middleware) {
	a.services[svcName] = svcFunc
}

// Option represent a function to modify pac app behavior
type AppOption func(*App)

// Middleware represent a middleware runs before handler. It accepts
// any number of options and return a handler to caller
type Middleware func(opts ...any) func(c *fiber.Ctx) error

// Service represent a package of related functions
type Service interface {
	Register(app *App, r *fiber.App)
}
