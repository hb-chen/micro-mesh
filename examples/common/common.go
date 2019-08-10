package common

import (
	grpc_zap "github.com/hb-go/grpc-contrib/log/zap"
	"github.com/hb-go/grpc-contrib/metadata"
	log_zap "github.com/hb-go/pkg/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func Logger(system string) error {
	// logger zap
	zapConf := zap.NewDevelopmentConfig()
	zapEncoderConf := zap.NewDevelopmentEncoderConfig()
	zapEncoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConf.EncoderConfig = zapEncoderConf
	logger, err := zapConf.Build(zap.AddCallerSkip(2))
	if err != nil {
		return err
	}

	grpc_zap.ReplaceGrpcLogger(logger)
	log_zap.ReplaceLogger(logger.WithOptions(zap.Fields(zap.String("system", system))))

	return nil
}

func GatewayMetadataOptions() []metadata.Option {
	return gatewayMetadataOptions()
}

func ClientInterceptors() []grpc.UnaryClientInterceptor {
	return clientInterceptors()
}

func ServerInterceptors() []grpc.UnaryServerInterceptor {
	return serverInterceptors()
}
