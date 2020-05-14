package grpcapi

import (
	"context"
	pb "hex-gopher/api/grpcapi/proto"
	"hex-gopher/app"
	"hex-gopher/repo/redisdb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	App app.GopherApp
}

func newServer(a app.GopherApp) *Server {
	return &Server{
		App: a,
	}
}

func StartServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to start tcp listener", err)
	}

	gopherdb, err := redisdb.NewRedisDB("localhost:6379", "")
	if err != nil {
		log.Fatal("failed to get gopher DB", err)
	}
	gopherApp := app.NewApp(gopherdb)
	gopherServer := newServer(gopherApp)

	grpcServer := grpc.NewServer()
	pb.RegisterGopherServiceServer(grpcServer, gopherServer)
	log.Fatal(grpcServer.Serve(lis))
}

func (s *Server) SaveGopher(ctx context.Context, req *pb.SaveGopherReq) (*pb.SaveGopherResp, error) {

	gopher := &app.Gopher{
		Id:   req.Id,
		Name: req.Name,
	}

	gopherId, err := s.App.SaveGopher(gopher)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to save gopher")
	}

	return &pb.SaveGopherResp{
		Id: gopherId,
	}, nil
}

func (s *Server) GetGopher(ctx context.Context, req *pb.GetGopherReq) (*pb.GetGopherResp, error) {

	gopher, err := s.App.GetGopher(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get gopher")
	}

	return &pb.GetGopherResp{
		Id:   gopher.Id,
		Name: gopher.Name,
	}, nil
}
