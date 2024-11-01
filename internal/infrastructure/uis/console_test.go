package uis_test

import (
	"errors"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/uis"
	"github.com/stretchr/testify/assert"
)

func Test_askCorrectData(t *testing.T) {
	const err = "err"

	printf := func(_ string, _ ...any) {}

	input1 := []any{"valid"}
	input2 := []any{"invalid", "valid"}
	input3 := []any{"err", "valid"}

	areValid := func(data ...any) bool {
		for _, d := range data {
			if *d.(*string) != "valid" {
				return false
			}
		}

		return true
	}

	type args struct {
		printf        func(format string, a ...any)
		read          func(data ...any) error
		areValid      func(data ...any) bool
		messageFormat string
		message       string
		errorMessage  string
		data          string
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "valid data",
			args: args{
				printf: printf,
				read: func(data ...any) error {
					testData := input1[0].(string)
					input1 = input1[1:]

					if testData == err {
						return errors.New("simulated error")
					}

					*data[0].(*string) = testData

					return nil
				},
				areValid:      areValid,
				messageFormat: "%s",
				message:       "",
				errorMessage:  "",
				data:          "",
			},
			expected: "valid",
		},
		{
			name: "invalid data, then valid",
			args: args{
				printf: printf,
				read: func(data ...any) error {
					testData := input2[0].(string)
					input2 = input2[1:]

					if testData == err {
						return errors.New("simulated error")
					}

					*data[0].(*string) = testData

					return nil
				},
				areValid:      areValid,
				messageFormat: "%s",
				message:       "",
				errorMessage:  "",
				data:          "",
			},
			expected: "valid",
		},
		{
			name: "err, then valid",
			args: args{
				printf: printf,
				read: func(data ...any) error {
					testData := input3[0].(string)
					input3 = input3[1:]

					if testData == err {
						return errors.New("simulated error")
					}

					*data[0].(*string) = testData

					return nil
				},
				areValid:      areValid,
				messageFormat: "%s",
				message:       "",
				errorMessage:  "",
				data:          "",
			},
			expected: "valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uis.AskCorrectData(
				tt.args.printf,
				tt.args.read,
				tt.args.areValid,
				tt.args.messageFormat,
				tt.args.message,
				tt.args.errorMessage,
				&tt.args.data,
			)

			assert.Equal(t, tt.expected, tt.args.data)
		})
	}
}
