package model

type TopicType struct {
	Model
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
}

func (f *TopicType) TableName() string {
	return "topic_type"
}