package models

// Note represents the task data structure
type Note struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"completed"`
	IsStarred  bool   `json:"starred"`
	Priority   int    `json:"priority"`
}
