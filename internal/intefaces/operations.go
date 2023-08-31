package interfaces

type Operations interface {
	LeftShift(value int) []byte
	RightShift(value int) []byte
}
