package grpc

import (
	"context"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"web-template/endpoints"
	"web-template/pb"
)

type (
	grpcServer struct {
		greeter grpctransport.Handler
	}
)

func NewGrpcServer(endpoints endpoints.Endpoints, logger log.Logger) *grpcServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{greeter: grpctransport.NewServer(endpoints.GreetingEndpoint, decodeGrpcGreetingRequest, encodeGrpcGreetingResponse, options...)}
}

func encodeGrpcGreetingResponse(ctx context.Context, grpcResp interface{}) (response interface{}, err error) {
	res := response.(endpoints.GreetingResponse)
	return &pb.GreetingResponse{Greeting: res.Greeting}, nil
}

func decodeGrpcGreetingRequest(ctx context.Context, grpcReq interface{}) (request interface{}, err error) {
	req := grpcReq.(*pb.GreetingRequest)
	return endpoints.GreetingRequest{Name: req.Name}, nil
}

func (s *grpcServer) Greeting(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	_, res, err := s.greeter.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GreetingResponse), nil
}
