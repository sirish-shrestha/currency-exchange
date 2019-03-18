package dbseeds

import (
	"io"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestReadEnvelope(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    XMLEnvelope
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadEnvelope(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadEnvelope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadEnvelope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRates(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRates(); got != tt.want {
				t.Errorf("GetRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveRate(t *testing.T) {
	type args struct {
		rates []*Rate
		db    *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveRate(tt.args.rates, tt.args.db)
		})
	}
}

func TestImportRates(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ImportRates(tt.args.db)
		})
	}
}
