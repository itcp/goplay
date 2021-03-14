package model

type Timg struct {
	ID        string    `json:"id" gorm:"column:id"`
	Tmid      string    `json:"tmid" gorm:"column:tmid"`
	Path      string    `json:"path" gorm:"column:path"`
}

func (f *Timg) TableName() string {
	return "topic_img"
}

func AddTopicImg(timg *Timg) (error) {
	timg.ID = GetID()

	// 入库
	err := db.Create(timg).Error
	return  err
}

func GetTopicImgByTid(topicID string)(timg []Timg, err error){
	err = db.Model(&Timg{}).Where("tmid = ?", topicID).Find(&timg).Error
	return
}