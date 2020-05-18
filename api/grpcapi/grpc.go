package grpcapi

import (
	"context"
	pb "hex-gopher/api/grpcapi/proto"
	"hex-gopher/app"
	"hex-gopher/config/env"
	"hex-gopher/repo/redisdb"
	"log"

	"github.com/ganeshdipdumbare/grpchelper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	App    app.GopherApp
	Server *grpchelper.Server
}

func NewServer(envVar *env.EnvVar) *Server {
	gopherdb, err := redisdb.NewRedisDB(envVar.RedisAddr, envVar.RedisPass)
	if err != nil {
		log.Fatal("failed to get gopher DB", err)
	}
	gopherApp := app.NewApp(gopherdb)
	s, err := grpchelper.NewServer(envVar.GrpcPort, []grpc.UnaryServerInterceptor{}, []grpc.StreamServerInterceptor{}, true)
	if err != nil {
		log.Fatal("error occurred while seting up grpc server")
	}

	srv := &Server{
		App:    gopherApp,
		Server: s,
	}
	pb.RegisterGopherServiceServer(s.GrpcServer, srv)

	return srv
}

func (s *Server) StartServer() {

	go func() {
		err := s.Server.Serve()
		if err != nil {
			log.Fatal("unable to start grpc server")
		}
	}()

	s.Server.AwaitTermination()
	log.Println("grpc server is stopped")
}

func (s *Server) StopServer() {
	s.Server.GrpcServer.GracefulStop()
	s.Server.Listner.Close()
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
