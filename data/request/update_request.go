package request

type UpdateRequest struct {
	ID      int
	Title   string `validate:"required min=1,max=100" json:"title"`
	Content string `json:"content"`
}
