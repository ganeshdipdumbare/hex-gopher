package mongodb

import (
	"hex-gopher/app"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func Test_getMongoClient(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should pass if correct uri is passed",
			args: args{
				uri: "mongodb://localhost:27017",
			},
			wantErr: false,
		},
		{
			name: "should return error if incorrect uri is passed",
			args: args{
				uri: "mongodb://localhost:27000",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getMongoClient(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMongoClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewMongoDB(t *testing.T) {
	type args struct {
		uri        string
		db         string
		collection string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should pass if correct inputs are passed",
			args: args{
				collection: "gophertestcollection",
				db:         "gophertestdb",
				uri:        "mongodb://localhost:27017",
			},
			wantErr: false,
		},
		{
			name: "should return error if incorrect inputs are passed",
			args: args{
				collection: "gophertestcollection",
				db:         "gophertestdb",
				uri:        "mongodb://localhost:27000",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMongoDB(tt.args.uri, tt.args.db, tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMongoDB_SaveGopher(t *testing.T) {
	mongoClient, err := getMongoClient("mongodb://localhost:27017")
	if err != nil {
		t.Fatal("failed to connect to mongo db")
	}

	type fields struct {
		client           *mongo.Client
		gopherDb         string
		gopherCollection string
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
			name: "should pass if correct inputs are passed",
			args: args{
				g: &app.Gopher{
					Id:   "0",
					Name: "gopher0",
				},
			},
			fields: fields{
				client:           mongoClient,
				gopherCollection: "gophertestcollection",
				gopherDb:         "gophertestdb",
			},
			want:    "0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MongoDB{
				client:           tt.fields.client,
				gopherDb:         tt.fields.gopherDb,
				gopherCollection: tt.fields.gopherCollection,
			}
			got, err := r.SaveGopher(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.SaveGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MongoDB.SaveGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDB_GetGopher(t *testing.T) {
	mongoDb, err := NewMongoDB("mongodb://localhost:27017", "gophertestdb", "gophertestcollection")
	if err != nil {
		t.Fatal("unable to connect to mongo db")
	}
	_, err = mongoDb.SaveGopher(&app.Gopher{
		Id:   "1",
		Name: "gopher1",
	})
	if err != nil {
		t.Fatal("unable to save gopher")
	}

	type fields struct {
		client           *mongo.Client
		gopherDb         string
		gopherCollection string
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
			name: "should pass for correct input",
			args: args{
				id: "1",
			},
			fields: fields{
				client:           mongoDb.client,
				gopherCollection: "gophertestcollection",
				gopherDb:         "gophertestdb",
			},
			want: &app.Gopher{
				Id:   "1",
				Name: "gopher1",
			},
			wantErr: false,
		},
		{
			name: "should return error for incorrect input",
			args: args{
				id: "-1",
			},
			fields: fields{
				client:           mongoDb.client,
				gopherCollection: "gophertestcollection",
				gopherDb:         "gophertestdb",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MongoDB{
				client:           tt.fields.client,
				gopherDb:         tt.fields.gopherDb,
				gopherCollection: tt.fields.gopherCollection,
			}
			got, err := r.GetGopher(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.GetGopher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoDB.GetGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}
