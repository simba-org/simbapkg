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
	"context"
	"fmt"
	"github.com/Bifang-Bird/simbapkg/balan"
	configs "github.com/Bifang-Bird/simbapkg/pkg/dbconfig"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type LoadBalanceHandler func(cfg *configs.Config) balan.LoadBalance

type InitGrpcHandler func(ctx context.Context) *grpc.Server

type BandingPortHandler func(cfg *config.HTTP, cancel context.CancelFunc) net.Listener

type Server struct {
	InitGrpcHandler    InitGrpcHandler
	BandingPortHandler BandingPortHandler
	LoadBalanceHandler LoadBalanceHandler
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

// PayChannelLoadBalance
//
//	@Description: 支付渠道负载均衡初始化
//	@param cfg
//	@return balan.LoadBalance
func InitLoadBalanceStrategy(cfg *configs.Config) balan.LoadBalance {
	//支付渠道相关的配置
	loadBalance := balan.LoadBalanceFactory(10)
	//支付渠道非指定时，需要初始化支付渠道的选举策略
	if cfg.LoadBalance.Specify {
		loadBalance = balan.LoadBalanceFactory(balan.LbConsistentHash)
		err := loadBalance.Add(cfg.LoadBalance.Channel)
		if err != nil {
			return nil
		}
	} else {
		if cfg.LoadBalance.SelectMode.Strategy > 2 {
			slog.Error("failed init payment channel,selectMode=", cfg.LoadBalance.SelectMode)
		} else {
			loadBalance = balan.LoadBalanceFactory(balan.LbType(cfg.LoadBalance.SelectMode.Strategy))
			for _, item := range cfg.LoadBalance.SelectMode.Weight {
				err := loadBalance.Add(item.Chan, item.Value)
				if err != nil {
					return nil
				}
			}
		}
	}
	slog.Info("负载策略初始化", cfg.LoadBalance.SelectMode)
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
		}))
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
