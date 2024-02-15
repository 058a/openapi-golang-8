package item

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewId(t *testing.T) {
	t.Parallel()

	type args struct {
		v uuid.UUID
	}
	newId := uuid.New()
	tests := []struct {
		name    string
		args    args
		want    Id
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", args{newId}, Id{newId}, false},
		{"nil", args{uuid.Nil}, Id{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewId(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestId_UUID(t *testing.T) {
	t.Parallel()

	type fields struct {
		value uuid.UUID
	}
	newId := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{"test", fields{newId}, newId},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := NewId(tt.fields.value)
			if err != nil {
				t.Fatal(err)
			}
			if got := v.UUID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Id.UUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestId_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		value uuid.UUID
	}
	newId := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test", fields{newId}, newId.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := NewId(tt.fields.value)
			if err != nil {
				t.Fatal(err)
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Id.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
