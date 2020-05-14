package redisdb

import (
	"hex-gopher/app"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

const (
	redis_addr = "localhost:6379"
	redis_db   = ""
)

func TestNewRedisDB(t *testing.T) {
	type args struct {
		addr string
		pwd  string
	}
	tests := []struct {
		name    string
		args    args
		want    *RedisDB
		wantErr bool
	}{
		{
			name: "should fail if db is unavailable",
			args: args{
				addr: "localhost:1234",
				pwd:  "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRedisDB(tt.args.addr, tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedisDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_SaveGopher(t *testing.T) {
	cl, err := getRedisClient(redis_addr, redis_db)
	if err != nil {
		t.Fatal("failed to get client for redis")
	}
	defer cl.Close()

	type fields struct {
		client *redis.Client
	}
	type args struct {
		g *app.Gopher
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should pass if correct gopher passed",
			fields: fields{
				client: cl,
			},
			args: args{
				g: &app.Gopher{
					Id:   "1",
					Name: "Gopher1",
				},
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "should fail if nil gopher passed",
			fields: fields{
				client: cl,
			},
			args: args{
				g: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisDB{
				client: tt.fields.client,
			}
			got, err := r.SaveGopher(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.SaveGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisDB.SaveGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_GetGopher(t *testing.T) {
	cl, err := getRedisClient(redis_addr, redis_db)
	if err != nil {
		t.Fatal("failed to get client for redis")
	}
	defer cl.Close()

	// add dummy record for testing
	cl.Set("1", "gopher1", 0)

	type fields struct {
		client *redis.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Gopher
		wantErr bool
	}{
		{
			name: "should pass if correct gopher id passed",
			fields: fields{
				client: cl,
			},
			args: args{
				id: "1",
			},
			want: &app.Gopher{
				Id:   "1",
				Name: "gopher1",
			},
			wantErr: false,
		},
		{
			name: "should fail if empty gopher id passed",
			fields: fields{
				client: cl,
			},
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisDB{
				client: tt.fields.client,
			}
			got, err := r.GetGopher(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.GetGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisDB.GetGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}
