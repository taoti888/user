package handler

import (
	"context"
	"github.com/taoti888/user/global"
	"github.com/taoti888/user/model"
	"github.com/taoti888/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PermissionsServer struct {
	proto.UnimplementedPermissionsServer
}

func (p *PermissionsServer) CreatePermissions(ctx context.Context, req *proto.CreatePermissionsRequest) (*proto.PermissionsInfoResponse, error) {
	// 判断权限是否已创建
	var permissions model.Permissions
	result := global.DB.Where("name = ?", req.Name).First(&permissions)
	if result.RowsAffected == 1 {
		global.LOG.Info("permissions already exists.")
		return nil, status.Errorf(codes.AlreadyExists, "permissions already exists.")
	}
	if result.Error != nil {
		global.LOG.Error("get permissions error: ", zap.Error(result.Error))
		return nil, result.Error
	}

	// 数据入库
	permissions.Name = req.Name
	permissions.Description = req.Description
	permissions.Resources = req.Resources
	permissions.Actions = req.Actions
	result = global.DB.Create(&permissions)
	if result.Error != nil {
		global.LOG.Error("create permissions error: ", zap.Error(result.Error))
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return permissions.PermissionsInfoResponse(), nil
}

func (p *PermissionsServer) GetPermissions(ctx context.Context, req *proto.GetPermissionsRequest) (*proto.PermissionsInfoResponse, error) {
	var permissions model.Permissions
	result := global.DB.First(&permissions, req.Id)
	if result.RowsAffected == 0 {
		global.LOG.Info("permissions not found.")
		return nil, status.Errorf(codes.NotFound, "permissions not found.")
	}
	if result.Error != nil {
		global.LOG.Error("get permissions error: ", zap.Error(result.Error))
		return nil, result.Error
	}
	return permissions.PermissionsInfoResponse(), nil
}

func (p *PermissionsServer) UpdatePermissions(ctx context.Context, req *proto.UpdatePermissionsRequest) (*proto.PermissionsInfoResponse, error) {
	var permissions model.Permissions
	result := global.DB.First(&permissions, req.Id)
	if result.RowsAffected == 0 {
		global.LOG.Info("permissions not found.")
		return nil, status.Errorf(codes.NotFound, "permissions not found.")
	}
	if result.Error != nil {
		global.LOG.Error("get permissions error: ", zap.Error(result.Error))
		return nil, result.Error
	}

	// 数据入库
	permissions.Description = req.Description
	permissions.Resources = req.Resources
	permissions.Actions = req.Actions
	result = global.DB.Save(&permissions)
	if result.Error != nil {
		global.LOG.Error("update permissions error: ", zap.Error(result.Error))
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return permissions.PermissionsInfoResponse(), nil
}

func (p *PermissionsServer) DeletePermissions(ctx context.Context, req *proto.DeletePermissionsRequest) (*emptypb.Empty, error) {
	// 判断待删除的权限是否存在
	var permissions model.Permissions
	result := global.DB.First(&permissions, req.Id)
	if result.RowsAffected == 0 {
		global.LOG.Info("permissions not found.")
		return nil, status.Errorf(codes.NotFound, "permissions not found.")
	}
	if result.Error != nil {
		global.LOG.Error("get permissions error: ", zap.Error(result.Error))
		return nil, result.Error
	}
	result = global.DB.Delete(&permissions)
	if result.Error != nil {
		global.LOG.Error("delete permissions error: ", zap.Error(result.Error))
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

func (p *PermissionsServer) AddResourceToPermissions(ctx context.Context, req *proto.AddResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Internal, "add resource to permissions")
}
