package signals

import (
	"reflect"
	"testing"
)

func TestCreateSignal(t *testing.T) {
	tests := []struct {
		name   string
		value  interface{}
		update interface{}
	}{
		{name: "int", value: 10, update: 20},
		{name: "string", value: "original", update: "updated"},
		{name: "float", value: 5.5, update: 10.10},
		{name: "bool", value: true, update: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get, set, _ := createSignal(tt.value)
			if got := get(); !reflect.DeepEqual(got, tt.value) {
				t.Errorf("createSignal() initial get = %v, want %v", got, tt.value)
			}

			set(tt.update)
			if got := get(); !reflect.DeepEqual(got, tt.update) {
				t.Errorf("createSignal() updated get = %v, want %v", got, tt.update)
			}
		})
	}
}

func TestCreateSignalSubscription(t *testing.T) {
	tests := []struct {
		name   string
		value  interface{}
		update interface{}
	}{
		{name: "int", value: 10, update: 20},
		{name: "string", value: "original", update: "updated"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, set, subscribe := createSignal(tt.value)

			changes := make([]interface{}, 0)
			subscribe(func(value interface{}) {
				changes = append(changes, value)
			})

			set(tt.update)

			if got := changes[0]; !reflect.DeepEqual(got, tt.update) {
				t.Errorf("createSignal() subscribed value = %v, want %v", got, tt.update)
			}
		})
	}
}
