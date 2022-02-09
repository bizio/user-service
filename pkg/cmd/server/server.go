package cmd

import (
	"flag"
	"fmt"

	"github.com/bizio/wa-srv-base/cache"
	"github.com/bizio/user-service/pkg/protocol/grpc"
	"github.com/bizio/user-service/pkg/protocol/rest"
	v1 "github.com/bizio/user-service/pkg/service/v1"
	"github.com/bizio/user-service/pkg/service/v1/cloudpubsub"
	"github.com/bizio/user-service/pkg/service/v1/data"
	"golang.org/x/net/context"
)

type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// Google Project Id, needed to connect to the datastore
	ProjectId string
}

func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.ProjectId, "project-id", "", "Google Project id")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("Invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("Invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	if len(cfg.ProjectId) == 0 {
		return fmt.Errorf("Invalid project id: '%s'", cfg.ProjectId)
	}

	cacheService := cache.NewRedisCacheService()
	v1Api := v1.NewUserServiceServer(
		data.NewDatastoreService(cfg.ProjectId, cacheService),
		cloudpubsub.NewMessageQueue(cfg.ProjectId),
		cacheService,
	)
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	// check queue messages
	go v1Api.CheckAlerts()

	return grpc.RunServer(ctx, v1Api, cfg.GRPCPort)
}
