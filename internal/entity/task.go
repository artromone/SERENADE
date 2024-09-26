package entity

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
	StartTime   string `json:"start"`
	EndTime     string `json:"end"`
	Priority    string `json:"priority"`
}
