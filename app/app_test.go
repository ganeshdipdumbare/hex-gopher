package app

import (
	"reflect"
	"testing"
)

func TestApp_SaveGopher(t *testing.T) {
	type args struct {
		g *Gopher
	}
	tests := []struct {
		name    string
		a       *App
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should pass with valid gopher",
			a:    &App{db: &DBSuccess{}},
			args: args{
				&Gopher{},
			},
			want:    "successid",
			wantErr: false,
		},
		{
			name: "should fail with failure in DB",
			a:    &App{db: &DBFailure{}},
			args: args{
				&Gopher{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.SaveGopher(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.SaveGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App.SaveGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_GetGopher(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		a       *App
		args    args
		want    *Gopher
		wantErr bool
	}{
		{
			name: "should pass with valid id",
			a: &App{
				db: &DBSuccess{},
			},
			args: args{
				id: "1",
			},
			want: &Gopher{
				Id:   "1",
				Name: "",
			},
			wantErr: false,
		},
		{
			name: "should fail with invalid id",
			a: &App{
				db: &DBFailure{},
			},
			args: args{
				id: "123",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetGopher(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.GetGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.GetGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewApp(t *testing.T) {
	type args struct {
		gdb GopherDB
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		{
			name: "should pass with valid DB",
			args: args{
				gdb: &DBSuccess{},
			},
			want: &App{
				db: &DBSuccess{},
			},
		},
		{
			name: "should pass with nil DB",
			args: args{
				gdb: nil,
			},
			want: &App{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(tt.args.gdb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}
