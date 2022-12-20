// Package health provies an in-process grpc_health_v1.HealthClient
// to a grpc_health_v1.HealthServer

package health

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// HealthClient is an in-process grpc_health_v1.HealthClient to a grpc_health_v1.HealthServer
type HealthClient struct {
	svc grpc_health_v1.HealthServer

	grpc_health_v1.HealthClient
}

func NewHealthClient(svc grpc_health_v1.HealthServer) *HealthClient {
	return &HealthClient{
		svc: svc,
	}
}

func (hc *HealthClient) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	return hc.svc.Check(ctx, in)
}
