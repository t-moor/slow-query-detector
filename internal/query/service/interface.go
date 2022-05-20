//go:generate mockery --name=Repository
package service

import (
	"context"

	"github.com/t-moor/slow-query-detector/internal/query/repository"
)

type Repository interface {
	FindQueries(ctx context.Context, criteria repository.FindQueriesCriteria) (repository.FindQueriesResults, error)
}
