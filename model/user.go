package model

import(
	"time"
	"goplay/public/common"
)

type User struct {
	ID        string       `gorm:"column:id"`
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
	hashPassword, _ := common.BcryptString(password)

	user = &User{
		ID: getsID(),
		Username:  username,
		Password:  hashPassword,
		Email:     email,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// 入库
	err = db.Create(user).Error

	return
}

func ExistUserByName(username string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&user)
	if user.ID == "" {
		return true
	}
	return false
}

func GetUserByName(username string)(user User, err error){
	err = db.Model(&User{}).Where("username = ?", username).First(&user).Error
	return
}