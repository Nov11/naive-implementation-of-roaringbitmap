package container

type TypeContainer int

const (
	BitmapContainerType      TypeContainer = 1
	SortedArrayContainerType TypeContainer = 2
	RunContainerType         TypeContainer = 3
)

type Container interface {
	exists(v uint16) bool
	add(v uint16) bool
	del(v uint16) bool
	convert(target TypeContainer) *Container
}
