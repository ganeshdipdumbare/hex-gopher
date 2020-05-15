package grpcapi

import (
	"context"
	pb "hex-gopher/api/grpcapi/proto"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	srv := NewServer()
	go srv.StartServer()
	time.Sleep(2 * time.Second)
	defer srv.StopServer()

	_, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Error("new server is not started properly")
	}
}

func TestServer_SaveGopher(t *testing.T) {
	srv := NewServer()
	go srv.StartServer()
	time.Sleep(2 * time.Second)
	defer srv.StopServer()

	type args struct {
		ctx context.Context
		req *pb.SaveGopherReq
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.SaveGopherResp
		wantErr bool
	}{
		{
			name: "should pass if valid gopher passed",
			args: args{
				ctx: context.Background(),
				req: &pb.SaveGopherReq{
					Id:   "1",
					Name: "gopher1",
				},
			},
			want: &pb.SaveGopherResp{
				Id: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.SaveGopher(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.SaveGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.SaveGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetGopher(t *testing.T) {
	srv := NewServer()
	go srv.StartServer()
	time.Sleep(2 * time.Second)
	defer srv.StopServer()

	_, err := srv.SaveGopher(context.Background(), &pb.SaveGopherReq{
		Id:   "1",
		Name: "gopher1",
	})
	if err != nil {
		t.Fatal("failed to save gopher")
	}

	type args struct {
		ctx context.Context
		req *pb.GetGopherReq
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetGopherResp
		wantErr bool
	}{
		{
			name: "should pass if valid gopher id is passed",
			args: args{
				ctx: context.Background(),
				req: &pb.GetGopherReq{
					Id: "1",
				},
			},
			want: &pb.GetGopherResp{
				Id:   "1",
				Name: "gopher1",
			},
			wantErr: false,
		},
		{
			name: "should pass if invalid gopher id is passed",
			args: args{
				ctx: context.Background(),
				req: &pb.GetGopherReq{
					Id: "0",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetGopher(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}
