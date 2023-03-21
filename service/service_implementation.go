package service

import (
	"assignTele/data/request"
	"assignTele/data/response"
	"assignTele/helper"
	"assignTele/models"
	"assignTele/utility"
	"context"
)

type BlogsServiceImpl struct {
	BlogService utility.BlogsRepo
}

func NewBlogServiceImpl(blog utility.BlogsRepo) BlogService {
	return &BlogsServiceImpl{BlogService: blog}
}

// Create implements BlogService
func (b *BlogsServiceImpl) Create(ctx context.Context, request request.BlogCreate) {
	blogs := models.Blogs{
		Title: request.Title,
	}
	b.BlogService.Save(ctx, blogs)
}

// Delete implements BlogService
func (b *BlogsServiceImpl) Delete(ctx context.Context, BlogID int) {
	blog, err := b.BlogService.FindByID(ctx, BlogID)
	helper.PanicIfError(err)
	b.BlogService.Delete(ctx, blog.ID)
}

// FindAll implements BlogService
func (b *BlogsServiceImpl) FindAll(ctx context.Context) []response.Response {
	blog := b.BlogService.FindAll(ctx)

	var blogsList []response.Response
	for _, value := range blog {
		blogs := response.Response{ID: value.ID, Title: value.Title, Content: value.Content}
		blogsList = append(blogsList, blogs)
	}

	return blogsList

}

// FindbyID implements BlogService
func (b *BlogsServiceImpl) FindbyID(ctx context.Context, BlogID int) response.Response {
	blog, err := b.BlogService.FindByID(ctx, BlogID)
	helper.PanicIfError(err)
	return response.Response(blog)
}

// Update implements BlogService
func (b *BlogsServiceImpl) Update(ctx context.Context, request request.UpdateRequest) {
	blog, err := b.BlogService.FindByID(ctx, request.ID)
	helper.PanicIfError(err)

	blog.Title = request.Title
	b.BlogService.Update(ctx, blog)
}
