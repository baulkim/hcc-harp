package server

import (
	"github.com/hcloud-classic/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hcc/harp/lib/config"
	"hcc/harp/lib/logger"
	"net"
	"strconv"
)

// Init : Initialize gRPC server
func Init() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(config.Grpc.Port)))
	if err != nil {
		logger.Logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHarpServer(s, &harpServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	logger.Logger.Println("Opening gRPC server on port " + strconv.Itoa(int(config.Grpc.Port)) + "...")
	if err := s.Serve(lis); err != nil {
		logger.Logger.Fatalf("failed to serve: %v", err)
	}
}
