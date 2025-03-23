package worker

type StorageAnalyse struct {
	Path string
}

func (receiver StorageAnalyse) Run(ctx *Context) error {
	return nil
}
