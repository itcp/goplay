package model

import(
	"time"
)

type Topic struct {
	ID        string       `json:"id" gorm:"column:id"`
	Ttid        int       `json:"ttid" gorm:"column:ttid"`
	Uid      string    `json:"uid" gorm:"column:uid"`
	Title      string    `json:"title" gorm:"column:title"`
	Keywords      string    `json:"keywords" gorm:"column:keywords"`
	Content      string    `json:"content" gorm:"column:content"`
	Mimg      string    `json:"m_img" gorm:"column:m_img"`
	Status     int        `json:"status" gorm:"column:status"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (f *Topic) TableName() string {
	return "topic_main"
}

func AddTopic(topic *Topic) (error) {
	topic.ID = getsID()
	topic.CreatedAt = time.Now()
	topic.UpdatedAt = time.Now()

	// 入库
	err := db.Create(topic).Error

	return err
}

func ExistTopicByTitle(title string) bool {
	var topic Topic
	db.Model(&Topic{}).Select("id").Where("title = ?", title).First(&topic)
	if topic.ID == "" {
		return true
	}
	return false
}

func GetTopicByID(topicID string)(topic Topic, err error){
	err = db.Model(&Topic{}).Where("id = ?", topicID).First(&topic).Error
	return
}