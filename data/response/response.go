package response

type Response struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
