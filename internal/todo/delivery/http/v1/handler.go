package v1

import (
	"github.com/Javanshir-SH/query-monitoring/internal/todo/entities"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TodoHandler struct {
	service Service
	logger  *zap.Logger
}

func NewTodoHandler(l *zap.Logger, svc Service) *TodoHandler {
	return &TodoHandler{logger: l, service: svc}
}

func (h *TodoHandler) Create(ctx *fiber.Ctx) error {
	task := entities.Task{}
	if err := ctx.BodyParser(&task); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	if err := task.Validate(); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	err := h.service.Create(task)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(400, "something went wrong"))
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (h *TodoHandler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	task, err := h.service.Get(id)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(500, "something went wrong"))
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (h *TodoHandler) Update(ctx *fiber.Ctx) error {
	task := entities.Task{}
	if err := ctx.BodyParser(&task); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	if err := task.Validate(); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	err := h.service.Update(task)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(500, "something went wrong"))
	}
	return ctx.Status(fiber.StatusOK).JSON([]byte{})
}

func (h *TodoHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	err = h.service.Delete(id)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(500, "something went wrong"))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
