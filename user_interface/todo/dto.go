package todo

type CreateRequest struct {
	Title       string `form:"title"`
	Category    string `form:"category"`
	Description string `form:"description"`
}

type UpdateRequest struct {
	Title       string `form:"title"`
	Category    string `form:"category"`
	Description string `form:"description"`
}
