package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `json:"name" validate:"required" min:"1" max:"200"`
}
