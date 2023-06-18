package model

type RolePermissions struct {
	RoleID        uint32 `gorm:"column:roleId;comment:'角色ID'"`
	PermissionsID uint32 `gorm:"column:permissionsId;comment:'权限ID'"`
}

func (r RolePermissions) TableName() string {
	return "u_role_permissions"
}
