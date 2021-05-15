package models

// Task represents the task data structure
type Task struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"completed"`
	IsStarred  bool   `json:"starred"`
	Priority   int    `json:"priority"`
}
