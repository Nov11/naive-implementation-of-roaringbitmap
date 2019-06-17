package container

//at most 2048 runs. 2 uint16 per run. first uint16 is start value. second is continuous count - 1
type RunContainer struct {
	value [4096]uint16
}
