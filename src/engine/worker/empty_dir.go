package worker

import (
	"BluebellSpace/util/fileutil"
)

type EmptyDir struct {
	IncludePath []string
}

func (receiver *EmptyDir) Run() *Context {
	list := make([]string, 0)
	for _, path := range receiver.IncludePath {
		err := fileutil.ListEmptyDir(path, &list)
		if err != nil {
			return nil
		}
	}
	return nil
}
