package text

import "io"

type FileMeta struct {
}

type Reader interface {
	GetFileSize(FileMeta) int64
	GetFileContent(FileMeta) []byte
	GetFileStream(FileMeta) io.Reader
}
