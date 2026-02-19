package component

import "reflect"

type Vector struct {
	data reflect.Value
}

func NewVector(sliceType reflect.Type) *Vector {
	if sliceType.Kind() != reflect.Slice {
		panic("[NewVector] the argument should be a slice type")
	}

	return &Vector{
		data: reflect.MakeSlice(sliceType, 0, 0),
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

func (v *Vector) Remove(index int) bool {
	length := v.Len()
	if index >= length {
		return false
	}

	lastIndex := length - 1

	if index != lastIndex {
		lastValue := v.data.Index(lastIndex)
		v.data.Index(index).Set(lastValue)
	}

	v.data = v.data.Slice(0, lastIndex)

	return true
}

func (v *Vector) Set(index int, value any) bool {
	if index >= v.Len() {
		return false
	}

	val := reflect.ValueOf(value)
	if !val.Type().AssignableTo(v.data.Type().Elem()) {
		panic("[VecterSet] value is wrong type for this vector")
	}

	v.data.Index(index).Set(val)
	return true
}

func (v *Vector) Get(index int) (any, bool) {
	if index >= v.Len() {
		return nil, false
	}

	return v.data.Index(index).Interface(), true
}
