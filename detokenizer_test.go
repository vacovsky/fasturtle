package main

import (
	"fmt"
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

func Test_detokenize(t *testing.T) {
	type args struct {
		input    []byte
		tokenMap []map[string][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{name: "file",
			args: args{
				input: []byte("\\\\File\\Share"),
				tokenMap: []map[string][]byte{
					map[string][]byte{},
				},
			},
			want: []byte(`\\File\Share`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detokenize(tt.args.input, tt.args.tokenMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detokenize() = %v, want %v", got, tt.want)
				fmt.Println(string(tt.args.input), string(tt.want))
			}
		})
	}
}
