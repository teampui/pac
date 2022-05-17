package pac

import (
	"github.com/gofiber/fiber/v2"
)

// NewApplication returns a Pac App
func NewApp(opts ...AppOption) *App {
	// create new pac app
	pacApp := App{
		HookCreated:      []AppOption{},
		HookBeforeCreate: []AppConfig{},
		Middlewares:      NewRegistry[Middleware](),
		Repositories:     NewRegistry[Repository](),
		Services:         NewRegistry[any](),
	}

	// apply options
	for _, opt := range opts {
		opt(&pacApp)
	}

	// then we init app
	fiberCfg := fiber.Config{
		DisableStartupMessage: true,
	}

	// apply before create hook
	for _, hook := range pacApp.HookBeforeCreate {
		hook(&fiberCfg)
	}

	// Pass the engine to the Views
	pacApp.fiber = fiber.New(fiberCfg)

	// apply created hook
	for _, hook := range pacApp.HookCreated {
		hook(&pacApp)
	}

	// clear everything after success called
	pacApp.HookCreated = nil
	pacApp.HookBeforeCreate = nil

	// return to caller
	return &pacApp
}

// App is Pac App which accept services and serves
type App struct {
	// fiber-related
	fiber *fiber.App
	// listen port
	port string
	// hooks
	HookCreated      []AppOption
	HookBeforeCreate []AppConfig
	// registry
	Middlewares  *Registry[Middleware]
	Repositories *Registry[Repository]
	Services     *Registry[any]
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

// Router returns internal Fiber App instance
func (a *App) Router() *fiber.App {
	return a.fiber
}

// Router returns internal Fiber App instance
func (a *App) Middleware(name string) Middleware {
	found := a.Middlewares.Get(name)

	if found == nil {
		return nil
	}

	return (*found)
}

// Repo[T] return repository and convert into user's request type
func Repo[T any](a *App, name string) *T {
	item := a.Repositories.Get(name)

	if item == nil {
		return nil
	}

	reqItem, ok := (*item).(T)

	if !ok {
		return nil
	}

	return &reqItem
}

// Svc[T] return service and convert into user's request type
func Svc[T any](a *App, name string) *T {
	item := a.Services.Get(name)

	if item == nil {
		return nil
	}

	reqItem, ok := (*item).(T)

	if !ok {
		return nil
	}

	return &reqItem
}

// Must[T] will check given variable, and force to assert it as type T
// panic user's error message if unsuccessful
func Must[T any](src any, error_msg string) T {
	if src == nil {
		panic(error_msg)
	}

	v, ok := src.(*T)

	if !ok {
		panic(error_msg)
	}

	if v == nil {
		panic(error_msg + "; maybe registered entity not fully implement the interface?")
	}

	return *v
}

// Option represent a function to modify pac app behavior
type AppConfig func(*fiber.Config)

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
