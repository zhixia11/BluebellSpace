package worker

import (
	"BluebellSpace/util/fileutil"
	"encoding/json"
	"fmt"
	"github.com/corona10/goimagehash"
	imgtype "github.com/shamsher31/goimgtype"
	"io/fs"
	"math/bits"
	"os"
	"path/filepath"
	"time"
)

type SimilarImage struct {
	Config  *SimilarImageConfig
	Context *Context
	Result  map[string][]Image
}

type SimilarImageConfig struct {
	Path []string `json:"path"`
}

type Image struct {
	Path    string                 `json:"path"`
	Hash    *goimagehash.ImageHash `json:"-"`
	Size    int64                  `json:"size"`
	ModTime time.Time              `json:"mod_time"`
	Width   int                    `json:"width"`
	Height  int                    `json:"height"`
}

func (i Image) MarshalJSON() ([]byte, error) {
	type Alias Image
	return json.Marshal(&struct {
		ModTime string `json:"mod_time"`
		*Alias
	}{
		ModTime: i.ModTime.Format("2006-01-02 15:04:05"),
		Alias:   (*Alias)(&i),
	})
}

type SimilarMessage struct {
	Hash  string `json:"hash"`
	Image *Image `json:"image"`
}

func NewSimilarImage(str string) (*SimilarImage, error) {
	config := &SimilarImageConfig{}
	err := json.Unmarshal([]byte(str), config)
	if err != nil {
		return nil, err
	}
	instance := &SimilarImage{
		Config:  config,
		Context: NewContext("similar"),
		Result:  make(map[string][]Image),
	}
	return instance, nil
}

func (receiver *SimilarImage) Run() {
	defer receiver.Context.Complete()
	receiver.Context.Info("开始扫描相似图片...")
	hashMap := make(map[uint8][]Image)
	paths := receiver.Config.Path
	if len(paths) == 0 {
		receiver.Context.Error("没有指定路径")
		return
	}
	// 遍历所有图片
	for _, root := range paths {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if receiver.Context.IsCancel() {
				return CancelError
			}
			if err != nil {
				//忽略错误
				receiver.Context.Error(fmt.Sprintf("访问路径 %s 时出错: %v", path, err))
				return nil
			}
			if info.IsDir() {
				// 忽略文件夹
				return nil
			}
			_, err = imgtype.Get(path)
			if err != nil {
				//非图片类型
				return nil
			}
			receiver.Context.Info(fmt.Sprintf("%s", path))
			image, err := readImageMeta(path)
			if err != nil {
				receiver.Context.Error(fmt.Sprintf("计算图片 %s 的感知哈希时出错: %v", path, err))
				return nil
			}
			cnt := uint8(bits.OnesCount64(image.Hash.GetHash()))
			hashMap[cnt] = append(hashMap[cnt], *image)
			//计算相似
			err = receiver.listSimilarImage(hashMap[cnt], *image)
			return nil
		})
	}
}

func (receiver *SimilarImage) listSimilarImage(list []Image, curr Image) error {
	for _, image := range list {
		if image.Path == curr.Path {
			continue
		}
		distance, err := image.Hash.Distance(curr.Hash)
		if err != nil {
			return err
		}
		if distance == 0 {
			hex := fmt.Sprintf("%x", image.Hash.GetHash())
			isFirst := receiver.Result[hex] == nil
			if isFirst {
				receiver.Result[hex] = append(receiver.Result[hex], image)
				receiver.Context.Data(&SimilarMessage{
					Hash:  hex,
					Image: &image,
				})
			}
			receiver.Result[hex] = append(receiver.Result[hex], curr)
			receiver.Context.Data(&SimilarMessage{
				Hash:  hex,
				Image: &curr,
			})
			return nil
		}
	}
	return nil
}

func readImageMeta(path string) (*Image, error) {
	img, err := fileutil.ReadImage(path)
	if err != nil {
		return nil, err
	}
	// 感知hash
	hash, err := fileutil.CalcImageHash(img)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	image := &Image{
		Path:    path,
		Hash:    hash,
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
		Width:   img.Bounds().Dx(),
		Height:  img.Bounds().Dy(),
	}
	return image, nil
}
