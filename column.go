package zurvan

type Column interface {
	Resize(length int)
	Len() int
	Remove(index int)
	Set(index int, value any)
	Get(index int) any
	AsSlice() any
	Push(value any)
}
