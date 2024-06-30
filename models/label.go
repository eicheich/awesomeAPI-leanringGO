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

func GetLabelsByUserID(userID int) ([]Label, error) {
	var labels []Label
	if err := DB.Where("user_id = ?", userID).Find(&labels).Error; err != nil {
		return nil, err
	}

	return labels, nil
}

func GetLabelByID(id int) (*Label, error) {
	var label Label
	if err := DB.Where("id = ?", id).First(&label).Error; err != nil {
		return nil, err
	}

	return &label, nil
}
