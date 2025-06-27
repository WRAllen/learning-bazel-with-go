package pkg

import (
	"fmt"
	"net"

	api "github.com/wrallen/sampleBazel6/api/hello/v1"
	"github.com/wrallen/sampleBazel6/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	cfg := zap.NewProductionConfig()
	// 自定义时间格式
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// 改成更常见的字段名
	cfg.EncoderConfig.TimeKey = "time"
	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Cannot initialize zap logger: %v", err))
	}
	defer logger.Sync()

	port := config.Global.Port

	// 启动一个监听器
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen on port %s: %v", port, err))
	}

	// grpc服务创建
	grpcServer := grpc.NewServer()
	// 注册hello相关接口
	api.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	// 可选开启 reflection，方便使用 grpcurl 进行调试
	// 例如通过 grpcurl -plaintext localhost:8080 list 查看当前的所有grpc服务
	reflection.Register(grpcServer)

	logger.Info("Starting gRPC server", zap.String("port", port))
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to start gRPC server: %v", err))
	}
}
