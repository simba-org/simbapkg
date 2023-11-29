/*
*

	@author: junwang
	@since: 2023/8/18
	@desc: //TODO

*
*/
package app

import (
	config "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"
	middleware2 "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/grpc/middleware"
	"context"
	"fmt"
	"github.com/Bifang-Bird/simbapkg/balan"
	configs "github.com/Bifang-Bird/simbapkg/pkg/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"strings"
	"time"
)

var (
	Logger *zap.Logger
)

type LoadBalanceHandler func(cfg *configs.LoadBalance) balan.LoadBalance

type InitGrpcHandler func(ctx context.Context) *grpc.Server

type BandingPortHandler func(cfg *config.HTTP, cancel context.CancelFunc) net.Listener

type InitLogHandler func(cfg *config.Log)

type Server struct {
	InitGrpcHandler    InitGrpcHandler
	BandingPortHandler BandingPortHandler
	LoadBalanceHandler LoadBalanceHandler
	InitLogHandler     InitLogHandler
}

func NewServer() *Server {
	srv := &Server{}
	return srv
}
func (s *Server) SetInitGrpcHandler(handle InitGrpcHandler) *Server {
	s.InitGrpcHandler = handle
	return s
}
func (s *Server) SetBandingPortHandler(handle BandingPortHandler) *Server {
	s.BandingPortHandler = handle
	return s
}
func (s *Server) SetLoadBalanceHandler(handle LoadBalanceHandler) *Server {
	s.LoadBalanceHandler = handle
	return s
}
func (s *Server) SetInitLogHandler(handle InitLogHandler) *Server {
	s.InitLogHandler = handle
	return s
}

// PayChannelLoadBalance
//
//	@Description: 支付渠道负载均衡初始化
//	@param cfg
//	@return balan.LoadBalance
func InitLoadBalanceStrategy(cfg *configs.LoadBalance) balan.LoadBalance {
	//支付渠道相关的配置
	loadBalance := balan.LoadBalanceFactory(10)
	//支付渠道非指定时，需要初始化支付渠道的选举策略
	if cfg.Specify {
		loadBalance = balan.LoadBalanceFactory(balan.LbConsistentHash)
		err := loadBalance.Add(cfg.Channel)
		if err != nil {
			return nil
		}
	} else {
		if cfg.SelectMode.Strategy > 2 {
			slog.Error("failed init payment channel,selectMode=", cfg.SelectMode)
		} else {
			loadBalance = balan.LoadBalanceFactory(balan.LbType(cfg.SelectMode.Strategy))
			for _, item := range cfg.SelectMode.Weight {
				err := loadBalance.Add(item.Chan, item.Value)
				if err != nil {
					return nil
				}
			}
		}
	}
	slog.Info("负载策略初始化", cfg.SelectMode)
	return loadBalance
}

func InitGrpcServer(ctx context.Context) *grpc.Server {
	server := grpc.NewServer(grpc.MaxConcurrentStreams(1000),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    20 * time.Second, // 每隔10秒ping一次客户端
			Timeout: 5 * time.Second,  // 等待5秒ping再次确认，则认为连接已死
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             20 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc_middleware.WithUnaryServerChain(
			middleware2.GrpcContext(),
			middleware2.GrpcRecover(),
			middleware2.GrpcLogger(),
		))
	slog.Info("GRPC SERVER 初始化完成")
	return server
}

func BandingPort(cfg *config.HTTP, cancel context.CancelFunc) net.Listener {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	network := "tcp"
	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}
	slog.Info("🌏 start server...", "address", address)
	return l
}

func InitLogger(cfg *config.Log) {
	var err error
	// 定义配置项
	zapConfig := zap.NewProductionConfig()
	var bugLevel = zap.InfoLevel
	if strings.EqualFold(cfg.Level, "debug") {
		bugLevel = zap.DebugLevel
	}
	if strings.EqualFold(cfg.Level, "warn") {
		bugLevel = zap.WarnLevel
	}
	if strings.EqualFold(cfg.Level, "error") {
		bugLevel = zap.ErrorLevel
	}
	if strings.EqualFold(cfg.Level, "info") {
		bugLevel = zap.InfoLevel
	}
	// 设置日志级别
	zapConfig.Level = zap.NewAtomicLevelAt(bugLevel) // 设置为 Debug 级别
	// 设置日志输出格式为 JSON 格式
	zapConfig.Encoding = "json"
	// 设置日志输出位置（可以是文件、标准输出等）
	zapConfig.OutputPaths = []string{"stdout"} // 输出到标准输出

	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "function",
		StacktraceKey:  "",
		SkipLineEnding: false,
		LineEnding:     "",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	Logger, err = zapConfig.Build()
	if err != nil {
		panic("Failed to initialize logger")
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
