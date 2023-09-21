package entity

import (
	"github.com/go-tron/example/config"
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/mysql"
	"github.com/go-tron/snowflake-id"
	"gorm.io/gorm"
)

var userIdWorker = snowflakeId.NewWithConfig15(config.Config)

type User struct {
	UserId     int64          `gorm:"primary_key" json:"userId"` //用户ID
	Name       string         `json:"name"`                      //姓名
	Avatar     string         `json:"avatar"`                    //头像
	Password   string         `json:"password"`                  //密码
	Phone      string         `json:"phone"`                     //手机
	Province   string         `json:"province"`                  //省
	City       string         `json:"city"`                      //市
	Email      string         `json:"email"`                     //邮件
	Gender     int            `json:"gender"`                    //性别 1:女 2:男
	Nation     string         `json:"nation"`                    //民族
	BirthDate  string         `json:"birthDate"`                 //出生日期
	Address    string         `json:"address"`                   //地址
	RegisterAt localTime.Time `json:"registerAt"`                //注册时间
	ChannelId  int            `json:"channelId"`                 //注册渠道
	Status     int            `json:"status"`                    //0禁用 1 正常 2 冻结
	Deleted    int            `json:"deleted,omitempty"`
}

func (*User) RecordNotFoundError() string {
	return "账户不存在"
}

func (*User) UniqueIndexErrors() []mysql.UniqueIndexError {
	return []mysql.UniqueIndexError{
		{IndexName: "user.name", ErrMsg: "账户名已存在"},
		{IndexName: "user.phone", ErrMsg: "手机号已存在"},
	}
}

func (*User) BeforeCreate(tx *gorm.DB) error {
	id, err := userIdWorker.NextId()
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("UserId", id)
	return nil
}
