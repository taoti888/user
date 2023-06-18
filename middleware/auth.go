package middleware

import (
	"context"
	"errors"
	"github.com/taoti888/user/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	apiKey string
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{apiKey: global.CONFIG.System.ApiKey}
}

// 仅鉴权apiKey
// Switch info.FullMethod 根据方法名称进行不同的鉴权逻辑
func (interceptor *AuthInterceptor) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// 调用鉴权函数
	if err := interceptor.authenticate(ctx); err != nil {
		global.LOG.Error("failed to auth enticate apiKey,error: ", zap.Error(err))
		return nil, err
	}
	return handler(ctx, req)
}

func (interceptor *AuthInterceptor) authenticate(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)
	if values := md["x-api-key"]; len(values) > 0 {
		if values[0] == interceptor.apiKey {
			return nil
		}
	}
	return errors.New("invalid api key")
}
