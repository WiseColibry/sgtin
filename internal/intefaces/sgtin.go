package interfaces

type Sgtin interface {
	Header(h []byte)
	Filter(f []byte)
	Partition(p []byte)
	Prefix(pr []byte)
	Ireference(ir []byte)
	Serial(sr []byte)
}
