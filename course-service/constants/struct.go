package constants

// UpdateContent data object to update lesson content
type UpdateContent struct {
	Action       string   `json:"action"`
	Version      string   `json:"version"`
	TimeRequired int      `json:"time_required"`
	Content      struct{} `json:"content"`
}
