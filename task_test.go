package ncmlu_api

import "testing"

func TestNcmluTask(t *testing.T) {
	type args struct {
		phone  string
		passwd string
		code   int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{
			phone:  "",
			passwd: "",
			code:   86,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NcmluTask(tt.args.phone, tt.args.passwd, tt.args.code)
		})
	}
}
