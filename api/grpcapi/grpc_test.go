package grpcapi

import (
	"context"
	pb "hex-gopher/api/grpcapi/proto"
	"hex-gopher/app"
	"reflect"
	"testing"
)

func Test_newServer(t *testing.T) {
	type args struct {
		a app.GopherApp
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "should pass if valid app passed",
			args: args{
				a: &app.App{},
			},
			want: &Server{
				App: &app.App{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newServer(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer()
		})
	}
}

func TestServer_SaveGopher(t *testing.T) {
	type fields struct {
		App app.GopherApp
	}
	type args struct {
		ctx context.Context
		req *pb.SaveGopherReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SaveGopherResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				App: tt.fields.App,
			}
			got, err := s.SaveGopher(tt.args.ctx, tt.args.req)
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
	type fields struct {
		App app.GopherApp
	}
	type args struct {
		ctx context.Context
		req *pb.GetGopherReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetGopherResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				App: tt.fields.App,
			}
			got, err := s.GetGopher(tt.args.ctx, tt.args.req)
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
