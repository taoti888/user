package model

import "time"

type User struct {
	BaseModel
	Username string    `json:"username" gorm:"column:username;unique;type:varchar(20);comment:'用户名称'"`
	Password string    `json:"password" gorm:"column:password;type:varchar(20);comment:'密码'"`
	Nickname string    `json:"nickname" gorm:"column:nickname;type:varchar(20);comment:'用户昵称'"`
	Email    string    `json:"email" gorm:"column:email;type:varchar(20);comment:'用户邮箱'"`
	Phone    string    `json:"phone" gorm:"column:phone;type:varchar(20);comment:'用户手机号码'"`
	Birthday time.Time `json:"birthday" gorm:"column:birthday;type:datetime;comment:'用户生日'"`
	RoleID   []int32   `json:"roleID" gorm:"column:roleID;type:integer;comment:'关联的角色ID列表'"`
}

func (User) TableName() string {
	return "u_user"
}
