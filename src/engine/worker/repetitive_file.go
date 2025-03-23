package worker

import (
	"BluebellSpace/util/fileutil"
	"fmt"
)

// RepetitiveFile 重复文件扫描
type RepetitiveFile struct {
	IncludePath []string
	IncludeExt  []string
}

// Run 运行
func (receiver *RepetitiveFile) Run(ctx *Context) error {
	sizeMap := make(map[int64][]string)
	hashMap := make(map[string][]string)
	// 扫描文件大小
	for _, path := range receiver.IncludePath {
		err := fileutil.ScanSize(path, receiver.IncludeExt, sizeMap)
		if err != nil {
			return err
		}
	}
	// 计算文件Hash
	for _, paths := range sizeMap {
		if len(paths) <= 1 {
			continue
		}
		for _, path := range paths {
			hash, err := fileutil.CalcFileHash(path)
			if err != nil {
				return err
			}
			if hashMap[hash] == nil {
				hashMap[hash] = make([]string, 0)
			}
			hashMap[hash] = append(hashMap[hash], path)
		}
	}
	for _, paths := range hashMap {
		if len(paths) <= 1 {
			continue
		}
		fmt.Println(paths)
	}
	return nil
}
