package repository

import (
	"context"
	"gorm.io/gorm"
)

// QueryAnalytics repository
type QueryAnalytics struct {
	db *gorm.DB
}

// NewQueryAnalytics constructor for QueryAnalytics
func NewQueryAnalytics(db *gorm.DB) *QueryAnalytics {
	return &QueryAnalytics{db: db}
}

func (s QueryAnalytics) FindQueries(ctx context.Context, criteria FindQueriesCriteria) (FindQueriesResults, error) {
	col := make(FindQueriesResults, 0)

	q := s.db.WithContext(ctx)

	if criteria.Command != "" {
		q = q.Where("starts_with(lower(query), lower(?))", criteria.Command)
	}

	switch criteria.Sort {
	case "fast-to-slow":
		q = q.Order("max_exec_time ASC")
	case "slow-to-fast":
		q = q.Order("max_exec_time DESC")
	default:
		return nil, ErrInvalidSort
	}

	q = q.Limit(criteria.PerPage).Offset((criteria.Page - 1) * criteria.PerPage)

	if err := q.Find(&col).Error; err != nil {
		return nil, err
	}

	return col, nil
}
