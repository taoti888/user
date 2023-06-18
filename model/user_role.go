package model

// SysUserAuthority 是 sysUser 和 sysAuthority 的连接表
type UserRole struct {
	UserId uint32 `gorm:"column:userId;comment:'用户ID'"`
	RoleID uint32 `gorm:"column:roleId;comment:'角色ID'"`
}

func (u *UserRole) TableName() string {
	return "u_user_role"
}
