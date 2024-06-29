package models

type Label struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
	UserID int `json:"user_id"`
}

func (label *Label) TableName() string {
	return "labels"
}
