package ui

import "testing"

func Test_askCorrectData(t *testing.T) {
	type args struct {
		printf        func(format string, a ...any) error
		read          func(data ...any) error
		messageFormat string
		message       string
		errorMessage  string
		areValid      func(data ...any) bool
		data          []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			askCorrectData(tt.args.printf, tt.args.read, tt.args.messageFormat, tt.args.message, tt.args.errorMessage, tt.args.areValid, tt.args.data...)
		})
	}
}
