package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/t-moor/slow-query-detector/internal/query/repository"
	"github.com/t-moor/slow-query-detector/internal/query/service/dto"
)

type QueryAnalytics struct {
	logger *zap.Logger
	repo   Repository
}

func NewQueryAnalytics(l *zap.Logger, repo Repository) *QueryAnalytics {
	return &QueryAnalytics{
		logger: l,
		repo:   repo,
	}
}

func (s *QueryAnalytics) FindQueries(ctx context.Context, input dto.FindQueriesInput) (dto.FindQueriesOutput, error) {

	criteria := repository.FindQueriesCriteria{
		Command: input.Command,
		Sort:    input.Sort,
		Page:    input.Page,
		PerPage: input.PerPage,
	}

	results, err := s.repo.FindQueries(ctx, criteria)
	if err != nil {
		return dto.FindQueriesOutput{}, err
	}

	output := make(dto.FindQueriesOutput, 0, len(results))
	for _, result := range results {
		output = append(output, &dto.QueryInfo{
			QueryID:       result.QueryID,
			Query:         result.Query,
			ExecutionTime: result.ExecutionTime,
		})
	}

	return output, nil
}
