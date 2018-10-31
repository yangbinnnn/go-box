package file

import (
	"io"
)

type File interface {
	Put(name string, input io.Reader) (int64, error)
	Get(name string, output io.Writer) (int64, error)
}
