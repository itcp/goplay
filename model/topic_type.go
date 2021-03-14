package model

type TopicType struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
}

func (f *TopicType) TableName() string {
	return "topic_type"
}

func GetTopicTypeList()(TopicTypeList []TopicType, err error) {
	err = db.Model(&TopicType{}).Find(&TopicTypeList).Error
	return
}

func GetTopicTypeByID(typeID string)(typelist TopicType, err error){
	err = db.Model(&TopicType{}).Where("id = ?", typeID).First(&typelist).Error
	return
}