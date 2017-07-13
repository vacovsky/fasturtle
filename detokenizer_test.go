package main

import (
	"reflect"
	"testing"
)

func Test_mapKeyPairs(t *testing.T) {
	type args struct {
		input  [][]byte
		buffer []string
	}
	tests := []struct {
		name string
		args args
		want []map[string][]byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapKeyPairs(tt.args.input, tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapKeyPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_outputToFile(t *testing.T) {
	type args struct {
		path string
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputToFile(tt.args.path, tt.args.data)
		})
	}
}
