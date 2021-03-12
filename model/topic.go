package model

import(
	"time"
)

type Topic struct {
	ID        string       `json:"id" gorm:"column:id"`
	Ttid        string       `json:"ttid" gorm:"column:ttid"`
	Uid      string    `json:"uid" gorm:"column:uid"`
	Title      string    `json:"title" gorm:"column:title"`
	Keywords      string    `json:"keywords" gorm:"column:keywords"`
	content      string    `json:"content" gorm:"column:content"`
	Mimg      string    `json:"m_img" gorm:"column:m_img"`
	Status     int        `json:"status" gorm:"column:status"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (f *Topic) TableName() string {
	return "topic_main"
}

