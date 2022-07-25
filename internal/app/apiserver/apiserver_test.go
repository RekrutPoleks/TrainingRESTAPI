package apiserver

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func TestAPIServer_Start(t *testing.T) {
	type fields struct {
		config *Config
		logger *logrus.Logger
		router *mux.Router
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &APIServer{
				config: tt.fields.config,
				logger: tt.fields.logger,
				router: tt.fields.router,
			}
			if err := s.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want *APIServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
