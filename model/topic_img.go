package model

import(
	"time"
)

type Timg struct {
	ID        string    `json:"id" gorm:"column:id"`
	Tmid      string    `json:"tmid" gorm:"column:tmid"`
	Path      string    `json:"path" gorm:"column:path"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (f *Timg) TableName() string {
	return "topic_img"
}