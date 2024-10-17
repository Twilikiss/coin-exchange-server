// Package tools
// @Author twilikiss 2024/5/2 13:56:56
package tools

import (
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
)

// Default 为传入的值设置默认值
func Default(data any) {
	// 这里我们借助反射处理
	t := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	if t.Kind() != reflect.Pointer {
		logx.Errorf("传入类型为[%s], 要求传递pointer类型", t.Kind())
		return
	}

	typeEle := t.Elem()
	valueEle := value.Elem()
	for i := 0; i < typeEle.NumField(); i++ {
		field := typeEle.Field(i)
		value := valueEle.Field(i)

		// 也可以通过在model的各个字段的tag上加上default
		// field.Tag.Get("default")
		kind := field.Type.Kind()

		// 简单支持几种常见的类型
		if kind == reflect.Int {
			value.Set(defaultInt())
		}
		if kind == reflect.Int32 {
			value.Set(defaultInt32())
		}
		if kind == reflect.Int64 {
			value.Set(defaultInt64())
		}
		if kind == reflect.String {
			value.Set(defaultString())
		}
		if kind == reflect.Float64 {
			value.Set(defaultFloat64())
		}
		if kind == reflect.Float32 {
			value.Set(defaultFloat32())
		}
	}
}

func defaultString() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

func defaultInt() reflect.Value {
	var i int = 0
	return reflect.ValueOf(i)
}

func defaultInt32() reflect.Value {
	var i int32 = 0
	return reflect.ValueOf(i)
}
func defaultInt64() reflect.Value {
	var i int64 = 0
	return reflect.ValueOf(i)
}

func defaultFloat64() reflect.Value {
	var i float64 = 0
	return reflect.ValueOf(i)
}
func defaultFloat32() reflect.Value {
	var i float32 = 0
	return reflect.ValueOf(i)
}
