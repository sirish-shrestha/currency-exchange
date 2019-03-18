package model

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestDBMigrate(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DBMigrate(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBMigrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
