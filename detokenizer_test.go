package main

import (
	"reflect"
	"testing"
)

func Test_mapKeyPairs(t *testing.T) {
	type args struct {
		path   string
		buffer string
	}
	tests := []struct {
		name string
		args args
		want map[string][]byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapKeyPairs(tt.args.path, tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapKeyPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detokenize(t *testing.T) {
	type args struct {
		input    []byte
		tokenMap map[string][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detokenize(tt.args.input, tt.args.tokenMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadFile() = %v, want %v", got, tt.want)
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

func Test_outputToStdout(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Does it print?",
			args: args{
				data: []byte("Does it print?"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputToStdout(tt.args.data)
		})
	}
}
