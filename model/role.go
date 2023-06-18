package model

type Role struct {
	BaseModel
	Name          string  `json:"name" gorm:"column:name;unique;type:varchar(20);comment:'角色名称'"`
	Description   string  `json:"description" gorm:"column:description;type:varchar(255);comment:'角色描述'"`
	PermissionsID []int32 `json:"permissionsID" gorm:"column:permissionsID;type:text;comment:'角色关联的权限列表id'"`
}

func (Role) TableName() string {
	return "u_role"
}
