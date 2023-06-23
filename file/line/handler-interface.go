package line

type Handler interface {
	Handle(line []byte)
}
