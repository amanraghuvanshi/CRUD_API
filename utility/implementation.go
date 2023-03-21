package utility

import (
	"assignTele/helper"
	"assignTele/models"
	"context"
	"database/sql"
	"errors"
)

type BlogsRepoImplementaion struct {
	Db *sql.DB
}

func NewBlogsRepo(Db *sql.DB) BlogsRepo {
	return &BlogsRepoImplementaion{Db: Db}
}

// delete
func (b *BlogsRepoImplementaion) Delete(ctx context.Context, blogID int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	SQL := "DELETE from blogs WHERE id =$1"
	_, anyErr := tx.ExecContext(ctx, SQL, blogID)
	helper.PanicIfError(anyErr)
}

// findALL
func (b *BlogsRepoImplementaion) FindAll(ctx context.Context) []models.Blogs {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	SQL := "SELECT * FROM blogs"
	result, anyErr := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(anyErr)
	defer result.Close()

	var blogs []models.Blogs
	for result.Next() {
		blog := models.Blogs{}
		err := result.Scan(&blog.ID, &blog.Title, &blog.Content)
		helper.PanicIfError(err)
		blogs = append(blogs, blog)
	}

	return blogs

}

// findbyID
func (b *BlogsRepoImplementaion) FindByID(ctx context.Context, blogID int) (models.Blogs, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	SQL := "SELECT * FROM blogs WHERE id=$1"
	result, anyErr := tx.QueryContext(ctx, SQL, blogID)
	helper.PanicIfError(anyErr)
	defer result.Close()

	blogs := models.Blogs{}

	if result.Next() {
		err := result.Scan(&blogs.ID, &blogs.Title, &blogs.Content)
		helper.PanicIfError(err)
		return blogs, nil
	} else {
		return blogs, errors.New("INVALID ID, BOOK NOT FOUND")
	}

}

// Save
func (b *BlogsRepoImplementaion) Save(ctx context.Context, blogs models.Blogs) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	SQL := "INSERT INTO blogs(title) value($1)"
	_, anyErr := tx.ExecContext(ctx, SQL, blogs.Title)
	helper.PanicIfError(anyErr)
}

func (b *BlogsRepoImplementaion) Update(ctx context.Context, blogs models.Blogs) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	SQL := "UPDATE blogs SET title=$1,content=$2 WHERE id=$3"

	_, anyErr := tx.ExecContext(ctx, SQL, blogs.Title, blogs.Content, blogs.ID)
	helper.PanicIfError(anyErr)

}
