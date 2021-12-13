package config

import "reflect"

type Value struct {
	value reflect.Value
}

func NewValue(value interface{}) *Value {
	return &Value{
		value: reflect.ValueOf(value),
	}
}

func (i *Value) IsValid() bool {
	return !i.value.IsNil() || !i.value.IsValid()
}

func (i *Value) GetString() string {
	if i.value.Kind() != reflect.String {
		return ""
	}

	return i.value.String()
}

func (i *Value) GetInteger() int {
	if i.value.Kind() != reflect.Int {
		return 0
	}

	return int(i.value.Int())
}

func (i *Value) GetBool() bool {
	if i.value.Kind() != reflect.Bool {
		return false
	}

	return i.value.Bool()
}
