//go:generate mockery --name=Service
package handler

import (
	"context"

	"github.com/t-moor/slow-query-detector/internal/query/service/dto"
)

type Service interface {
	FindQueries(ctx context.Context, input dto.FindQueriesInput) (dto.FindQueriesOutput, error)
}
