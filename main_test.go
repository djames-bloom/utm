package utm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRouteType tests the routeType function with various inputs.
func TestRouteType(t *testing.T) {
	tests := []struct {
		name          string
		param         string
		value         any
		defaultValue  any
		expectedOpt   Option
		expectedError string
	}{
		{
			name:          "uint option",
			param:         "uintParam",
			value:         uint(42),
			defaultValue:  uint(0),
			expectedOpt:   Uint("uintParam", 42, 0),
			expectedError: "",
		},
		{
			name:          "string option",
			param:         "stringParam",
			value:         "example",
			defaultValue:  "",
			expectedOpt:   String("stringParam", "example", ""),
			expectedError: "",
		},
		{
			name:          "int option",
			param:         "intParam",
			value:         10,
			defaultValue:  0,
			expectedOpt:   Int("intParam", 10, 0),
			expectedError: "",
		},
		{
			name:          "string array option",
			param:         "arrayParam",
			value:         []string{"one", "two", "three"},
			defaultValue:  []string{},
			expectedOpt:   StringArray{Param: "arrayParam", Values: []string{"one", "two", "three"}},
			expectedError: "",
		},
		{
			name:          "unsupported type",
			param:         "floatParam",
			value:         float64(3.14),
			defaultValue:  float64(0),
			expectedOpt:   nil,
			expectedError: "unsupported type: float64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opt Option
			var err error

			switch val := tt.value.(type) {
			case string:
				opt, err = routeType(tt.param, val, tt.defaultValue.(string))
			case uint:
				opt, err = routeType(tt.param, val, tt.defaultValue.(uint))
			case int:
				opt, err = routeType(tt.param, val, tt.defaultValue.(int))
			case []string:
				opt, err = routeType(tt.param, val, tt.defaultValue.([]string))
			default:
				err = fmt.Errorf("unsupported type: %T", val)
			}

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOpt.OptionString(), opt.OptionString())
			}
		})
	}
}

// TestProcessOptions tests the ProcessOptions function.
func TestProcessOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []Option
		expected string
	}{
		{
			name: "multiple options",
			options: []Option{
				Uint("uintParam", 42, 0),
				String("stringParam", "example", ""),
				Int("intParam", 10, 0),
				StringArray{Param: "arrayParam", Values: []string{"one", "two", "three"}},
			},
			expected: "uintParam 42 stringParam example intParam 10 arrayParam one arrayParam two arrayParam three ",
		},
		{
			name:     "no options",
			options:  []Option{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProcessOptions(tt.options)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkProcessOptions(b *testing.B) {
	benchmarks := []struct {
		name    string
		numOpts int
	}{
		{"10 options", 10},
		{"100 options", 100},
		{"1000 options", 1000},
		{"10000 options", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			options := make([]Option, bm.numOpts)
			for i := 0; i < bm.numOpts; i++ {
				options[i] = String("param"+strconv.Itoa(i), "value"+strconv.Itoa(i), "")
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = ProcessOptions(options)
			}
		})
	}
}

func BenchmarkProcessOptionsStringArray(b *testing.B) {
	benchmarks := []struct {
		name      string
		numOpts   int
		arraySize int
	}{
		{"10 options, 10 elements", 10, 10},
		{"100 options, 10 elements", 100, 10},
		{"1000 options, 10 elements", 1000, 10},
		{"10000 options, 10 elements", 10000, 10},
		{"10 options, 100 elements", 10, 100},
		{"100 options, 100 elements", 100, 100},
		{"1000 options, 100 elements", 1000, 100},
		{"10000 options, 100 elements", 10000, 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			options := make([]Option, bm.numOpts)
			for i := 0; i < bm.numOpts; i++ {
				values := make([]string, bm.arraySize)
				for j := 0; j < bm.arraySize; j++ {
					values[j] = "value" + strconv.Itoa(j)
				}
				options[i] = StringArray{Param: "param" + strconv.Itoa(i), Values: values}
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = ProcessOptions(options)
			}
		})
	}
}
