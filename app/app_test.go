package app

import (
	"testing"
	"zumata-currency-exchange/config"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func TestApp_Initialize(t *testing.T) {
	type fields struct {
		Router *chi.Mux
		DB     *gorm.DB
	}
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.Initialize(tt.args.config)
		})
	}
}

func TestApp_setRouters(t *testing.T) {
	type fields struct {
		Router *chi.Mux
		DB     *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.setRouters()
		})
	}
}

func TestApp_Run(t *testing.T) {
	type fields struct {
		Router *chi.Mux
		DB     *gorm.DB
	}
	type args struct {
		host string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.Run(tt.args.host)
		})
	}
}
