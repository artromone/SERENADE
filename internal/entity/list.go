package entity

type List struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"task"`
}
