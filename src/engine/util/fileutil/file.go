package fileutil

import (
  "crypto/sha1"
  "encoding/hex"
  "fmt"
  "io"
  "os"
  "path/filepath"
  "slices"
)

// ScanSize 扫描文件夹(含子文件)大小
func ScanSize(root string, includeExt []string, sizeMap map[int64][]string) error {
  // 使用 filepath.Walk 遍历文件夹
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Printf("访问路径 %q 时出错: %v\n", path, err)
      return err
    }
    // 如果是文件，打印文件路径和大小
    if info.IsDir() {
      return nil
    }
    // 过滤文件后缀
    if len(includeExt) > 0 {
      ext := filepath.Ext(path)
      if !slices.Contains(includeExt, ext) {
        return nil
      }
    }
    if sizeMap[info.Size()] == nil {
      sizeMap[info.Size()] = make([]string, 0)
    }
    sizeMap[info.Size()] = append(sizeMap[info.Size()], path)
    return nil
  })
  return err
}

// CalcFileHash 计算文件哈希值
func CalcFileHash(path string) (string, error) {
  // 打开文件
  file, err := os.Open(path)
  if err != nil {
    return "", err
  }
  defer file.Close()
  // 创建 SHA1 哈希计算器
  hash := sha1.New()
  // 将文件内容写入哈希计算器
  if _, err := io.Copy(hash, file); err != nil {
    return "", err
  }
  hashInBytes := hash.Sum(nil)
  hashString := hex.EncodeToString(hashInBytes)
  return hashString, nil
}

// ListEmptyDir 列出目录下的所有空文件夹
func ListEmptyDir(root string, list *[]string) error {
  // 获取路径信息
  info, err := os.Stat(root)
  if err != nil {
    return err
  }

  // 如果是文件，直接返回
  if !info.IsDir() {
    return err
  }
  // 判断当前文件夹是否为空
  isEmptyDir, err := IsEmptyDir(root)
  if err != nil {
    return err
  }
  if isEmptyDir {
    // 如果当前文件夹为空，将当前文件夹添加到结果列表
    *list = append(*list, root)
  } else {
    // 读取目录内容
    entries, err := os.ReadDir(root)
    if err != nil {
      return err
    }
    // 递归子目录
    for _, entry := range entries {
      if entry.IsDir() {
        // 如果是子目录，递归检查子目录
        err := ListEmptyDir(filepath.Join(root, entry.Name()), list)
        if err != nil {
          return err
        }
      }
    }
  }
  return nil
}

// IsEmptyDir 判断文件夹是否为空
// 如果该文件夹下的文件夹全是空，那么该文件夹也会是空
func IsEmptyDir(root string) (bool, error) {
  // 获取路径信息
  info, err := os.Stat(root)
  if err != nil {
    return false, err
  }

  // 如果是文件，直接返回 false
  if !info.IsDir() {
    return false, nil
  }

  // 读取目录内容
  entries, err := os.ReadDir(root)
  if err != nil {
    return false, err
  }

  // 遍历目录内容
  for _, entry := range entries {
    if entry.IsDir() {
      // 如果是子目录，递归检查子目录是否为空
      isEmpty, err := IsEmptyDir(filepath.Join(root, entry.Name()))
      if err != nil {
        return false, err
      }
      if !isEmpty {
        return false, nil
      }
    } else {
      // 如果有文件，目录不为空
      return false, nil
    }
  }
  // 如果没有文件或子目录，或者所有子目录都是空的，目录为空
  return true, nil
}

type FileInfo struct {
  Name    string      `json:"name"`
  Size    int64       `json:"size"`
  Mode    os.FileMode `json:"mode"`
  ModTime string      `json:"mod_time" `
  IsDir   bool        `json:"is_dir"`
}

// GetFileInfo 获取文件信息
func GetFileInfo(path string) (*FileInfo, error) {
  fileInfo, err := os.Stat(path)
  if err != nil {
    return nil, err
  }
  result := &FileInfo{
    Name:    fileInfo.Name(),
    Size:    fileInfo.Size(),
    Mode:    fileInfo.Mode(),
    ModTime: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
    IsDir:   fileInfo.IsDir(),
  }
  return result, nil
}
