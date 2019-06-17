package container

//8KB at most. elements are uint16 integers
type SortedArray struct {
	value [4096]uint16
}
