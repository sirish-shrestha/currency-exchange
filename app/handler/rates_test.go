package handler

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func TestRateRoutes(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RateRoutes(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RateRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLatestRates(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLatestRates(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatestRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRatesByDate(t *testing.T) {
	type args struct {
		db   *gorm.DB
		date string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRatesByDate(tt.args.db, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRatesByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzeRates(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnalyzeRates(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzeRates() = %v, want %v", got, tt.want)
			}
		})
	}
}
