package data

import "context"

type Models struct {
	Post interface {
		Create(ctx context.Context, post *Post) error

		GetByID(ctx context.Context, id string) (*Post, error)

		GetByUserID(ctx context.Context, id string, skip, limit int64) (*[]Post, error)

		GetByCategory(ctx context.Context, category string, skip, limit int64) (*[]Post, error)

		GetByTags(ctx context.Context, tags []string, skip, limit int64) (*[]Post, error)

		Get(ctx context.Context, skip, limit int64) (*[]Post, error)

		UpdateByID(ctx context.Context, id string, post *Post) error

		DeleteByID(ctx context.Context, id string) error

		DeleteByUserID(ctx context.Context, id string) error
	}
}
