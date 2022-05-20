package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/t-moor/slow-query-detector/internal/query/handler/mocks"
	"github.com/t-moor/slow-query-detector/internal/query/service/dto"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryAnalytics_FindQueries(t *testing.T) {
	l, _ := zap.NewProduction()
	v := validator.New()

	tests := []struct {
		name       string
		handler    *QueryAnalytics
		route      string
		statusCode int
	}{
		{
			name:       "invalid page param",
			handler:    NewQueryAnalytics(l, v, &mocks.Service{}),
			route:      "/queries?command=select&sort=slow-to-fast&page=b&per-page=2",
			statusCode: 400,
		},
		{
			name:       "invalid per-page param",
			handler:    NewQueryAnalytics(l, v, &mocks.Service{}),
			route:      "/queries?command=select&sort=slow-to-fast&page=1&per-page=a",
			statusCode: 400,
		},
		{
			name:       "request validation error",
			handler:    NewQueryAnalytics(l, v, &mocks.Service{}),
			route:      "/queries?command=select&sort=asc&page=1&per-page=2",
			statusCode: 400,
		},
		{
			name: "service error",
			handler: NewQueryAnalytics(l, v, func() *mocks.Service {
				m := &mocks.Service{}
				m.On("FindQueries", mock.Anything, mock.Anything).Return(dto.FindQueriesOutput{}, errors.New("some error"))

				return m
			}()),
			route:      "/queries?command=select&sort=slow-to-fast&page=1&per-page=2",
			statusCode: 500,
		},
		{
			name: "success",
			handler: NewQueryAnalytics(l, v, func() *mocks.Service {
				m := &mocks.Service{}
				m.On("FindQueries", mock.Anything, mock.Anything).Return(dto.FindQueriesOutput{}, nil)

				return m
			}()),
			route:      "/queries?command=select&sort=slow-to-fast&page=1&per-page=2",
			statusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/queries", tt.handler.FindQueries)
			resp, err := app.Test(httptest.NewRequest(
				http.MethodGet,
				tt.route,
				nil,
			))

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}
