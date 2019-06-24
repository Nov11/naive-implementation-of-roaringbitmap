package container

type TypeContainer int

const (
	BitmapContainerType      TypeContainer = 1
	SortedArrayContainerType TypeContainer = 2
	RunContainerType         TypeContainer = 3
)

type Container interface {
	Exists(v uint16) bool
	Add(v uint16) bool
	Del(v uint16) bool
	convert(target TypeContainer) *Container
}
