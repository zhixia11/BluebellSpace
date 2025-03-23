package fileutil

import (
	"github.com/corona10/goimagehash"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// CalcImageHash 计算图片感知hash
func CalcImageHash(img image.Image) (*goimagehash.ImageHash, error) {
	hash, err := goimagehash.PerceptionHash(img)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// ReadImage 根据路径读取图片
func ReadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
