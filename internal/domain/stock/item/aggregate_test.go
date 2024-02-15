package item

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	type args struct {
		id   Id
		name Name
	}

	validId, err := NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	validName, err := NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		args args
		want *Aggregate
	}{
		{"test", args{validId, validName}, &Aggregate{validId, validName, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewAggregate(tt.args.id, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestoreAggregate(t *testing.T) {
	type args struct {
		id      Id
		name    Name
		deleted bool
	}

	validId, err := NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	validName, err := NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		args args
		want *Aggregate
	}{
		{"test", args{validId, validName, true}, &Aggregate{validId, validName, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreAggregate(tt.args.id, tt.args.name, tt.args.deleted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate_IsDeleted(t *testing.T) {
	type fields struct {
		Id      Id
		Name    Name
		deleted bool
	}

	validId, err := NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	validName, err := NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"active", fields{validId, validName, false}, false},
		{"deleted", fields{validId, validName, true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := RestoreAggregate(tt.fields.Id, tt.fields.Name, tt.fields.deleted)
			if got := a.IsDeleted(); got != tt.want {
				t.Errorf("Aggregate.IsDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate_Delete(t *testing.T) {
	type fields struct {
		Id   Id
		Name Name
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test", fields{Id{}, Name{}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAggregate(tt.fields.Id, tt.fields.Name)
			a.Delete()
			if got := a.IsDeleted(); got != tt.want {
				t.Errorf("Aggregate.IsDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
