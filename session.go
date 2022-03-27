package pac

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// SessionOperator is a toolkit to manipulate session within current context
type SessionOperator struct {
	ctxSession *session.Session
}

// NewSessionOperator returns a new SessionOperator to given context
func NewSessionOperator(c *fiber.Ctx) *SessionOperator {
	return &SessionOperator{
		ctxSession: c.Locals("session").(*session.Session),
	}
}

// SessionInjector is a middleware that automatically inject session into ctx.
// It will store session to ctx.Locals.
func SessionInjector(s *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": "cannot get session",
			})
		}

		c.Locals("session", sess)
		c.Next()

		if err := sess.Save(); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": "cannot update session",
			})
		}

		return nil
	}
}

// Destroy clear entire session data
func (s *SessionOperator) Destroy() error {
	return s.ctxSession.Destroy()
}

// Flash will returns values of requested variable and purge it
func (s *SessionOperator) Flash(var_name string, default_value interface{}) interface{} {
	// get flash message
	flash_var := s.ctxSession.Get(var_name)
	s.ctxSession.Delete(var_name)

	// return value
	if flash_var == nil {
		return default_value
	}

	return flash_var
}

// Get will returns values of requested variable
func (s *SessionOperator) Get(var_name string, default_value interface{}) interface{} {
	v := s.ctxSession.Get(var_name)

	if v == nil {
		return default_value
	}

	return v
}

// Set will put values into given variable
func (s *SessionOperator) Set(key string, value interface{}) {
	s.ctxSession.Set(key, value)
}
