package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/t-moor/slow-query-detector/internal/query/repository"
	"github.com/t-moor/slow-query-detector/internal/query/service/dto"
	"github.com/t-moor/slow-query-detector/internal/query/service/mocks"
	"go.uber.org/zap"
	"testing"
)

func TestQueryAnalytics_FindQueries(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name    string
		repo    Repository
		wantErr bool
	}{
		{
			name: "repo error",
			repo: func() Repository {
				m := &mocks.Repository{}
				m.On("FindQueries", mock.Anything, mock.Anything).Return(repository.FindQueriesResults{}, errors.New("db connection error"))

				return m
			}(),
			wantErr: true,
		},
		{
			name: "success",
			repo: func() Repository {
				m := &mocks.Repository{}
				m.On("FindQueries", mock.Anything, mock.Anything).Return(repository.FindQueriesResults{}, nil)

				return m
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewQueryAnalytics(l, tt.repo)

			_, err := svc.FindQueries(context.Background(), dto.FindQueriesInput{})
			if (err != nil) && !tt.wantErr {
				t.Fatal("didn't expect error")
			}

		})
	}
}
