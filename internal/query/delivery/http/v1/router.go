package v1

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterQueryMonitoringHandler(app *fiber.App, h *QueryMonitoringHandler) {
	app.Get("/queries", h.List)
}
