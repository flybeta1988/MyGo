package pkgOne

type ReadWriter interface {
	Read(buf []byte) (n int)
	Write(buf []byte) (n int)
}
