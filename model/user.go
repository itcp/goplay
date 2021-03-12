package model

import(
	"time"
)

type User struct {
	Model
	ID        string       `json:"id" gorm:"column:id"`
	Username      string    `json:"username" gorm:"column:username"`
	Password      string    `json:"password" gorm:"column:password"`
	Email      string    `json:"email" gorm:"column:email"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (f *User) TableName() string {
	return "user_main"
}


func AddUser(username, password, email string) (user *User, err error) {
	// 加密密码


	user = &User{
		ID: getsID(),
		Username:  username,
		Password:  password,
		Email:     email,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// 入库
	err = db.Create(user).Error

	return
}