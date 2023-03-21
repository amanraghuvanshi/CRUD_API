package utility

import (
	"assignTele/models"
	"context"
)

type BlogsRepo interface {
	Save(ctx context.Context, blogs models.Blogs)
	Update(ctx context.Context, blogs models.Blogs)
	Delete(ctx context.Context, blogID int)
	FindByID(ctx context.Context, blogID int) (models.Blogs, error)
	FindAll(ctx context.Context) []models.Blogs
}
