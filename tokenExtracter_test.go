package main

import (
	"reflect"
	"testing"
)

func Test_convertToJSON(t *testing.T) {
	type args struct {
		data   [][]byte
		buffer string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Ensure JSON converts",
			args: args{
				buffer: "HAHA",
				data: [][]byte{
					[]byte("HAHAtest123HAHA"),
				},
			},
			want: []byte(`{"test123":""}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToJSON(tt.args.data, tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractTokens(t *testing.T) {
	type args struct {
		input  []byte
		buffer string
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test extracting tokens",
			args: args{
				input:  []byte("bfuewiowrby8arwobyuv8fo@@@testkey@@@@@@TESTKEY2@@@bt34870fb78wpbu8pf"),
				buffer: "@@@",
			},
			want: [][]byte{[]byte("@@@testkey@@@"), []byte("@@@TESTKEY2@@@")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractTokens(tt.args.input, tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataBagsContents(t *testing.T) {

}

func Test_dataBagsList(t *testing.T) {
	data := collectDataBagJSON("_default", "connection_strings")
	// spew.Dump(data)
	mapKeyPairs([][]byte{data}, "__")
}
