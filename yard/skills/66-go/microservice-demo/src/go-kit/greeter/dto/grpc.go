package dto

import (
	"bbxyard.com/msvc/go-kit/greeter/ep"
	pb "bbxyard.com/msvc/pbs"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	greeter grpc.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC GreeterServer.
func NewGRPCServer(endpoints ep.Endpoints, logger log.Logger) pb.GreeterServer {
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{greeter: grpc.NewServer(
		endpoints.GreetingEndpoint,
		decodeGRPCGreetingRequest,
		encodeGRPCGreetingResponse,
		options...,
	),
	}
}

// Greeting implementation of the method of the GreeterService interface.
func (s *grpcServer) Greeting(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	_, res, err := s.greeter.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GreetingResponse), nil
}

// decodeGRPCGreetingRequest is a transport/grpc.DecodeRequestFunc that converts
// a gRPC greeting request to a user-domain greeting request.
func decodeGRPCGreetingRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GreetingRequest)
	return ep.GreetingRequest{Name: req.Name}, nil
}

// encodeGRPCGreetingResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain greeting response to a gRPC greeting response.
func encodeGRPCGreetingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ep.GreetingResponse)
	return &pb.GreetingResponse{Greeting: res.Greeting}, nil
}
