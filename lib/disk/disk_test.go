package disk

import (
	"reflect"
	"testing"
)

func TestDirSize(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirSize(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserList(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
