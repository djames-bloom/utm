package utm

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// HandledTypes is a type constraint for supported types.
type HandledTypes interface {
	string | uint | int | []string
}

// Option is an interface in Go that requires implementing the `OptionString()` method to return a string.
type Option interface {
	OptionString() string
}

// StringArray represents an option with a string parameter and an array of string values.
type StringArray struct {
	Param  string
	Values []string
}

// GenericOption represents an option with a generic value and associated functions.
type GenericOption[T any] struct {
	isDefault    func(T, T) bool
	toString     func(T) string
	Value        T
	defaultValue T
	param        string
}

// ProcessOptions concatenates the string representations of a slice of Option objects.
func ProcessOptions(options []Option) string {
	var buf bytes.Buffer
	for _, opt := range options {
		buf.WriteString(opt.OptionString())
	}
	return buf.String()
}

// NewGenericOption creates a new instance of a generic option with specified parameters.
func NewGenericOption[T any](param string, value T, defaultValue T, isDefault func(T, T) bool, toString func(T) string) *GenericOption[T] {
	return &GenericOption[T]{
		param:        param,
		Value:        value,
		defaultValue: defaultValue,
		isDefault:    isDefault,
		toString:     toString,
	}
}

// OptionString generates the schema-required formatted string representation of the option.
func (o *GenericOption[T]) OptionString() string {
	if !o.isDefault(o.Value, o.defaultValue) {
		return o.param + " " + o.toString(o.Value) + " "
	}
	return ""
}

// Uint creates a new generic option for a uint type with specified parameters.
func Uint(param string, value uint, defaultValue uint) *GenericOption[uint] {
	return NewGenericOption(param, value, defaultValue, func(a, b uint) bool { return a == b }, func(v uint) string { return strconv.FormatUint(uint64(v), 10) })
}

// String creates a new generic option with string values.
func String(param string, value string, defaultValue string) *GenericOption[string] {
	return NewGenericOption(param, value, defaultValue, func(a, b string) bool { return a == b }, func(v string) string { return v })
}

// Int creates a new generic option with integer values.
func Int(param string, value int, defaultValue int) *GenericOption[int] {
	return NewGenericOption(param, value, defaultValue, func(a, b int) bool { return a == b }, func(v int) string { return strconv.Itoa(v) })
}

// OptionString generates the formatted string representation of the StringArrayOption.
func (o StringArray) OptionString() string {
	var buf bytes.Buffer
	if len(o.Values) > 0 {
		for _, v := range o.Values {
			buf.WriteString(o.Param + " " + v + " ")
		}
	}
	return buf.String()
}

// routeType routes the value to the appropriate option type and returns an error for unsupported types.
func routeType[T HandledTypes](param string, value, defaultValue T) (Option, error) {
	// Use a type switch to handle each type explicitly.
	switch v := any(value).(type) {
	case string:
		return String(param, v, any(defaultValue).(string)), nil
	case uint:
		return Uint(param, v, any(defaultValue).(uint)), nil
	case int:
		return Int(param, v, any(defaultValue).(int)), nil
	case []string:
		return StringArray{Param: param, Values: v}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported type: %T", v))
	}
}
