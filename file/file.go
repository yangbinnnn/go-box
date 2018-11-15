package file

import (
	"go-box/common"
	"io"
)

type File interface {
	Put(name string, input io.Reader) (int64, error)
	Get(name string, output io.Writer) (int64, error)
	Size(name string) (int64, error)
}

var (
	FileStore File
)

func InitFile() {
	config := common.GlobalConfig
	switch config.File.Type {
	case "local":
		FileStore = NewLocal(config.File.Home)
		break
	case "s3":
		break
	default:
		FileStore = NewLocal(config.File.Home)
	}
}
