package dtos

type InputServices struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
