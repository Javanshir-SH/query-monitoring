package v1

import "github.com/gofiber/fiber/v2"

func RegisterTodoHandler(app *fiber.App, h *TodoHandler) {
	app.Post("/todos", h.Create)
	app.Get("/todos/:id", h.Get)
	app.Put("/todos", h.Update)
	app.Delete("/todos/:id", h.Delete)
}
