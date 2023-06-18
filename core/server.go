package core

import (
	"context"
	"fmt"
	"github.com/taoti888/user/global"
	"github.com/taoti888/user/handler"
	"github.com/taoti888/user/initialize"
	"github.com/taoti888/user/middleware"
	"github.com/taoti888/user/proto"
	"github.com/taoti888/user/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// HTTP健康检查
func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "healthy")
}

func setupHealthCheckEndpoint(ip string) {
	http.HandleFunc("/health", healthCheck)
	go http.ListenAndServe(fmt.Sprintf("%s:8081", ip), nil)
}

// 注册grpc服务
func registerServer(server *grpc.Server) {
	proto.RegisterPermissionsServer(server, &handler.PermissionsServer{})

	// GRPC健康检查
	//healthSrv := health.NewServer()
	//grpc_health_v1.RegisterHealthServer(server, healthSrv)
}

// 注册到consul
func registerServiceToConsul(ip, uuid string) {
	if err := initialize.NewConsulRegistrar(ip, uuid); err != nil {
		global.LOG.Error("initialize consul registrar failed,", zap.Error(err))
		os.Exit(1)
	}
	global.LOG.Info(fmt.Sprintf("Register to consul service %s:%d success!", global.CONFIG.Consul.Address, global.CONFIG.Consul.Port))
}

func RunWindowsServer() {
	// 获取ip地址
	netip, err := utils.BoundIP()
	if err != nil {
		global.LOG.Error("get bond ip address failed,error:", zap.Error(err))
		return
	}
	ip := netip.String()

	// 监听服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.CONFIG.System.Port))
	if err != nil {
		global.LOG.Error("failed to listen server,", zap.Error(err))
		return
	}

	// apiKey拦截器
	var interceptor = middleware.NewAuthInterceptor()

	// uuid 和 New server
	var uuid = utils.UUID()
	var server = grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryInterceptor))

	// 注册GRPC服务
	registerServer(server)
	global.LOG.Info(fmt.Sprintf("Register grpc server: %s success! UUID: %s", global.CONFIG.System.Name, uuid))

	// http模式的健康检查
	setupHealthCheckEndpoint(ip)

	// 注册到consul
	registerServiceToConsul(ip, uuid)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = server.Serve(listener); err != nil {
			global.LOG.Error("failed to start grpc server,", zap.Error(err))
		}
	}()

	global.LOG.Info(fmt.Sprintf("Grpc Server listening on %s:%d", ip, global.CONFIG.System.Port))

	// 优雅启动服务和关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	global.LOG.Info("Shutdown signal is received, ready to Deregister the service in consul...")

	// 设置超时以进行优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shutdownGraceful := make(chan struct{})
	go func() {
		defer close(shutdownGraceful)
		server.GracefulStop()
	}()

	select {
	case <-ctx.Done():
		global.LOG.Warn("Timeout during GracefulStop, forcing server to stop...")
		server.Stop()
	case <-shutdownGraceful:
		global.LOG.Info("gRPC server graceful shutdown completed")
	}

	// 带超时注销服务
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		err = initialize.Deregister(uuid)
	}()

	select {
	case <-ctx.Done():
		global.LOG.Warn("Timeout during service deregistration, forcing exit")
	case <-done:
		if err != nil {
			global.LOG.Error("failed to deregister,service is not exists or had deregistered!" + err.Error())
			return
		}
		global.LOG.Info("Deregistration complete!")
	}

	wg.Wait()
}
