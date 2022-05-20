package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/t-moor/slow-query-detector/internal/query/service/dto"
	"github.com/t-moor/slow-query-detector/pkg/rest"
)

const (
	SortSlowToFast = "slow-to-fast"
	SortFastToSlow = "fast-to-slow"
)

var defaultPage = "1"
var defaultPerPage = "10"

type QueryAnalytics struct {
	logger    *zap.Logger
	validator *validator.Validate
	svc       Service
}

func NewQueryAnalytics(l *zap.Logger, validator *validator.Validate, svc Service) *QueryAnalytics {
	return &QueryAnalytics{logger: l, validator: validator, svc: svc}
}

func (h *QueryAnalytics) FindQueries(ctx *fiber.Ctx) error {

	// get query command
	queryCommand := ctx.Query("command")

	// get sort type
	sort := ctx.Query("sort", SortSlowToFast)

	// get page param
	pageStr := ctx.Query("page", defaultPage)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err.Error())

		return ctx.
			Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestResponse("invalid \"page\" param, should be an integer", err))
	}

	// get per-page param
	perPageStr := ctx.Query("per-page", defaultPerPage)
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		h.logger.Error(err.Error())

		return ctx.
			Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestResponse("invalid \"per-page\" param, should be an integer", err))
	}

	req := rest.FindQueriesRequest{
		Command: queryCommand,
		Sort:    sort,
		Page:    page,
		PerPage: perPage,
	}

	if err := h.validator.Struct(&req); err != nil {
		h.logger.Error(err.Error())

		return ctx.
			Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestResponse("invalid input data", err))
	}

	input := dto.FindQueriesInput{
		Command: queryCommand,
		Sort:    sort,
		Page:    page,
		PerPage: perPage,
	}

	// call service
	output, err := h.svc.FindQueries(ctx.Context(), input)
	if err != nil {
		h.logger.Error(err.Error())

		return ctx.
			Status(http.StatusInternalServerError).
			JSON(rest.NewInternalServerErrorResponse("something went wrong, please contact administrator", err))
	}

	// prepare response
	resp := make(rest.FindQueriesResponse, 0, len(output))
	for _, row := range output {
		resp = append(resp, rest.QueryInfo{
			QueryID:       row.QueryID,
			Query:         row.Query,
			ExecutionTime: row.ExecutionTime,
		})
	}

	// prepare response
	return ctx.JSON(resp)
}
