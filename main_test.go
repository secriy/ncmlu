package main

import "testing"

func Test_execTask(t *testing.T) {
	type args struct {
		phone   string
		passwd  string
		expired string
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{
			expired: "2021-09-03",
		}},
		{"", args{
			expired: "2021-09-04",
		}},
		{"", args{
			expired: "2021-09-02",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execTask(tt.args.phone, tt.args.passwd, tt.args.expired)
		})
	}
}
