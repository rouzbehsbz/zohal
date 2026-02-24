package zurvan

import "reflect"

type Vector struct {
	data reflect.Value
}

func NewVector(elemType reflect.Type) *Vector {
	return &Vector{
		data: reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0),
	}
}

func (v *Vector) Resize(length int) {
	current := v.data.Len()

	if length <= current {
		v.data = v.data.Slice(0, length)
		return
	}

	elemType := v.data.Type().Elem()
	extra := reflect.MakeSlice(
		reflect.SliceOf(elemType),
		length-current,
		length-current,
	)

	v.data = reflect.AppendSlice(v.data, extra)
}

func (v *Vector) Len() int {
	return v.data.Len()
}

func (v *Vector) Remove(index int) {
	length := v.Len()
	if index >= length {
		return
	}

	lastIndex := length - 1

	if index != lastIndex {
		lastValue := v.data.Index(lastIndex)
		v.data.Index(index).Set(lastValue)
	}

	v.data = v.data.Slice(0, lastIndex)
}

func (v *Vector) Set(index int, value any) {
	if index >= v.Len() {
		return
	}

	val := reflect.ValueOf(value)
	v.data.Index(index).Set(val)
}

func (v *Vector) Get(index int) any {
	if index >= v.Len() {
		return nil
	}

	return v.data.Index(index).Interface()
}

func (v *Vector) AsSlice() any {
	return v.data.Interface()
}

func (v *Vector) Push(value any) {
	val := reflect.ValueOf(value)

	v.data = reflect.Append(v.data, val)
}
