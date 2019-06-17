package container

type CONTAINTER_TYPE int

const (
	BITMAP CONTAINTER_TYPE = 1
	ARRAY  CONTAINTER_TYPE = 2
	RUN    CONTAINTER_TYPE = 3
)

type Container interface {
	exists(v uint16) bool
	add(v uint16) bool
	remove(v uint16) bool
	convert(target CONTAINTER_TYPE) Container
}
