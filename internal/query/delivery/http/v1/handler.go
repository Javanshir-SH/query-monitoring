package v1

import (
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type QueryMonitoringHandler struct {
	logger  *zap.Logger
	service Service
}

func NewQueryMonitoringHandler(l *zap.Logger, svc Service) *QueryMonitoringHandler {
	return &QueryMonitoringHandler{logger: l, service: svc}
}

func (h *QueryMonitoringHandler) List(ctx *fiber.Ctx) error {
	// Get request param
	params := FilterQueryParam{}

	err := ctx.QueryParser(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	// check & set defaults for a page and per_page params
	params.checkAndSetDefaults()

	// Validate
	err = params.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, err.Error()))
	}

	// Call Service method
	in := service.FilterOptionDto{
		Command: params.Type,
		Sort:    params.Sort,
		Page:    params.Page,
		PerPage: params.PerPage,
	}

	queries, totalCount, pageCount, err := h.service.List(ctx.Context(), in)
	if err != nil {
		h.logger.Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(400, internalServerError))
	}

	// build response
	payload := make([]Payload, 0, len(queries))
	for _, row := range queries {
		payload = append(payload, Payload{
			ID:                row.QueryID,
			Statement:         row.Query,
			MeanExecutionTime: row.MeanExecutionTime,
			MaxExecutionTime:  row.MaxExecutionTime,
		})
	}

	resp := Response{
		Page:       params.Page,
		PerPage:    params.PerPage,
		PageCount:  pageCount,
		TotalCount: totalCount,
		Data:       payload,
	}

	// prepare response
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
