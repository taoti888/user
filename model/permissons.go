package model

import "github.com/taoti888/user/proto"

type Permissions struct {
	BaseModel
	Name        string   `json:"name" gorm:"column:name;unique;type:varchar(20);comment:'权限名称'"`
	Description string   `json:"description" gorm:"column:description;type:varchar(255);comment:'权限描述'"`
	Resources   []string `json:"resources" gorm:"column:resources;type:text;comment:'权限关联的资源项'"`
	Actions     []string `json:"actions" gorm:"column:actions;type:varchar(100);comment:'权限关联的动作'"`
}

func (p *Permissions) TableName() string {
	return "u_permissions"
}

func (p *Permissions) PermissionsInfoResponse() *proto.PermissionsInfoResponse {
	return &proto.PermissionsInfoResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Resources:   p.Resources,
		Actions:     p.Actions,
	}
}
