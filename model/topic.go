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

func AddTopic(topic *Topic) (Tid string, err error) {
	topic.ID = GetID()
	topic.CreatedAt = time.Now()
	topic.UpdatedAt = time.Now()

	// 入库
	err = db.Create(topic).Error

	return topic.ID,err
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

// 返回分页数据
func GetPageTopic(tid, page int)(topicList []Topic, err error){
	size := PAGE_SIZE
	if page > 1 {
		page = page - 1
	}else if page ==1 {
		page = 0
	}
	offsetNum := page * size 
	if tid == 0 {
		err = db.Model(&Topic{}).Where("status = ?", 2).Order("create_at desc").Offset(offsetNum).Limit(size).Find(&topicList).Error
	}else{
		err = db.Model(&Topic{}).Where("status = ?", 2).Where("ttid = ?", tid).Order("create_at desc").Offset(offsetNum).Limit(size).Find(&topicList).Error
	}
	return
}

func GetPassTopic()(topicList []Topic, err error){
	err = db.Model(&Topic{}).Where("status = ?", 1).Find(&topicList).Error
	return
}

func PassTopic(ptid string, pstatus int)( err error){
	err = db.Model(&Topic{}).Where("id = ?", ptid).Update("status", pstatus).Error
	return
}