package models

import (
	"database/sql"
)

type User struct {
	Id            int
	Mobile        string         `gorm:"type:char(11);index;unique;not null;" json:"mobile,omitempty"` //手机号,加索引，唯一，不为空
	Name          string         `gorm:"type:varchar(12);"`                                            //用户昵称，3-12个字符
	Desc          string         `gorm:"type:varchar(100);"`
	Sex           int            `gorm:"type:tinyint(1);default:0;"`
	Age           string         `gorm:"type:char(13);"` //用户年龄，存储的是时间戳字符串
	Avatar        string         `gorm:"type:varchar(255);"`
	FollowNum     int            `gorm:"default:0;"`
	FansNum       int            `gorm:"default:0;"`
	State         int            `gorm:"type:tinyint(1);default:0;" json:"-"` //用户状态，比如=1账号冻结，=2不允许聊天之类的,默认=0,为json时不返回
	WeixinOpenid  sql.NullString `gorm:"unique;index;" json:"-"`              //微信openeid，唯一，加索引，json不返回，注意类型使用sql.NullString。因为设置了唯一，而string默认为""，所以会有冲突，一下同理
	WeixinUnionid sql.NullString `gorm:"unique;" json:"-"`                    //微信unionid，唯一，json不返回
	QqOpenid      sql.NullString `gorm:"unique;index;" json:"-"`              //QQopenid,唯一，加索引，json不返回
}

func FindUserByMobile(mobile string) (User, error) {
	var user User
	err := db.Where("mobile=?", mobile).First(&user).Error
	return user, err

}

func CreateUser(mobile string) (User, error) {
	var user User
	if mobile != "" {
		user.Mobile = mobile
		err := db.Create(&user).Error
		return user, err
	}
	return user, nil
}
