package models

type CreateSubmissionRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Track       string `json:"track" validate:"required"`
	GithubLink  string `json:"github_link"`
	FigmaLink   string `json:"figma_link`
	OtherLink   string `json:"other_link"`
}

type UpdateSubmissionRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Track       string `json:"track" validate:"required"`
	GithubLink  string `json:"github_link"`
	FigmaLink   string `json:"figma_link"`
	OtherLink   string `json:"other_link"`
}
