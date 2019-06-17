package container

//8KB at most. elements are uint16 integers
type SortedArrayContainer struct {
	value [4096]uint16
}
