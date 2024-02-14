package item

import (
	"reflect"
	"testing"
)

func TestNewName(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{"test", args{"test"}, Name{"test"}, false},
		{"", args{""}, Name{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_String(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test", fields{"test"}, "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Name{
				string: tt.fields.string,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Name.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
