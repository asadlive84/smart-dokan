package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryLoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Before RPC: %s", info.FullMethod)

	resp, err := handler(ctx, req)

	log.Printf("After RPC: %s, error=%v", info.FullMethod, err)

	return resp, err
}
