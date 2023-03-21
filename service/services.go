package service

import (
	"assignTele/data/request"
	"assignTele/data/response"
	"context"
)

type BlogService interface {
	Create(ctx context.Context, request request.BlogCreate)
	Update(ctx context.Context, request request.UpdateRequest)
	Delete(ctx context.Context, BlogID int)
	FindbyID(ctx context.Context, BlogID int) response.Response
	FindAll(ctx context.Context) []response.Response
}
