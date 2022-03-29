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
	hookCreated  []AppOption
	middlewares  map[string]Middleware
	repositories map[string]Repository
}

// Add will add service to pac app, callin' its `Register()` to inject routes
// and register service functions to app.
func (a *App) Add(svc Service) {
	svc.Register(a)
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

// Middleware return function by its register name, return nil if not found
func (a *App) Middleware(svcName string) Middleware {
	svc, ok := a.middlewares[svcName]

	if !ok {
		return nil
	}

	return svc
}

// RegisterMiddleware register middleware function to pac app, get by call `Middleware()`
func (a *App) RegisterMiddleware(svcName string, svcFunc Middleware) {
	a.middlewares[svcName] = svcFunc
}

// Repository return requested repository from registry, return nil if not found
func (a *App) Repository(repoName string) Repository {
	svc, ok := a.repositories[repoName]

	if !ok {
		return nil
	}

	return svc
}

// RegisterRepository put repo into registry
func (a *App) RegisterRepository(repoName string, repo Repository) {
	a.repositories[repoName] = repo
}

// Router returns internal Fiber App instance
func (a *App) Router() *fiber.App {
	return a.fiber
}

// Option represent a function to modify pac app behavior
type AppOption func(*App)

// Middleware represent a middleware runs before handler. It accepts
// any number of options and return a handler to caller
type Middleware func(opts ...any) func(c *fiber.Ctx) error

// Repository is actual same as service
type Repository Service

// Service represent a package of related functions
type Service interface {
	Register(app *App)
}
