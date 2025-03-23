package service

import (
	"BluebellSpace/util/fileutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type MultiFileReq struct {
	Paths []string `json:"paths"`
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	req := &MultiFileReq{}
	err := c.BindJSON(req)
	if err != nil || len(req.Paths) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	errs := make([]string, 0)
	for _, path := range req.Paths {
		err = os.Remove(path)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": errs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func FileInfo(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	info, err := fileutil.GetFileInfo(path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": info,
	})
}
