package pkgTwo

type IStream interface {
	Read(buf []byte) (n int)
	Write(buf []byte) (n int)
}
