package model

import (
	"fmt"
	"log"
	"math/rand"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"github.com/BurntSushi/toml"
)

var db *gorm.DB

const PAGE_SIZE int = 9

type Model struct {
	gorm.Model
}
type tomlConfig struct {
	DB      database `toml:"database"`
}
type database struct {
	DriverName  string `toml:"driver_name"`
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
	DatabaseName string `toml:"database_name"`
	Username string
	Password string
}

var config tomlConfig

func init() {
	if _, err := toml.DecodeFile("conf/dev/config.toml", &config); err != nil {
		return 
	}
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = config.DB.DriverName
	dbName = config.DB.DatabaseName
	user = config.DB.Username
	password = config.DB.Password
	host = config.DB.Server
	tablePrefix = ""

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})
}

func CloseDB() {
	defer db.Close()
}

func GetDb() *gorm.DB {
	return db
}

// 自增
func Increment(table string, id int, colum string) error {
	//pcolum := colum + " + 1"
	return db.Exec("update ? set files_num = files_num + 1 where id = 1", "bbs_post").Error
}

func getsID() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 9; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
