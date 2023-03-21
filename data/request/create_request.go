package request

type BlogCreate struct {
	Title string `validate:"required min=1,max=100" json:"title"`
}
