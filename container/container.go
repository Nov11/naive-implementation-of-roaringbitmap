package container

type CONTAINTER_TYPE int

const (
	BitmapContainerType      CONTAINTER_TYPE = 1
	SortedArrayContainerType CONTAINTER_TYPE = 2
	RunContainerType         CONTAINTER_TYPE = 3
)

type Container interface {
	exists(v uint16) bool
	add(v uint16) bool
	remove(v uint16) bool
	convert(target CONTAINTER_TYPE) *Container
}
