package service

import (
	"BluebellSpace/worker"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var workerNameMap = make(map[string]*worker.Context)
var workerContextMap = make(map[string]*worker.Context)

type WorkerConf struct {
	Name   string                 `json:"name"`
	Config map[string]interface{} `json:"config"`
}

// SubmitWorker 提交任务
func SubmitWorker(c *gin.Context) {
	conf := &WorkerConf{}
	err := c.BindJSON(conf)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	ctx := workerNameMap[conf.Name]
	if ctx == nil {
		str, _ := json.Marshal(conf.Config)
		instance, err := worker.NewSimilarImage(string(str))
		if err != nil {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": "模块配置错误",
			})
			return
		}
		ctx = instance.Context
		workerNameMap[conf.Name] = ctx
		workerContextMap[ctx.ID] = ctx
		go instance.Run()
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": ctx.ID,
	})
}

// WorkerProgress 获取进度
func WorkerProgress(c *gin.Context) {
	ctxId := c.Query("id")
	if workerContextMap[ctxId] == nil {
		message, _ := json.Marshal(gin.H{
			"type": "error",
			"data": "任务不存在",
		})
		c.String(200, string(message))
		return
	}
	// 设置响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	ctx := workerContextMap[ctxId]
	defer func() {
		delete(workerNameMap, ctx.Name)
		delete(workerContextMap, ctxId)
	}()
	// 创建一个通道用于发送消息
	ch := make(chan string)
	go func() {
		index := 0
		for !ctx.IsCancel() {
			for ; index < len(ctx.Logs); index++ {
				message, _ := json.Marshal(ctx.Logs[index])
				ch <- string(message)
			}
		}
		ch <- "complete"
		// 关闭通道
		close(ch)
	}()

	// 监听客户端断开连接
	clientGone := c.Request.Context().Done()

	// 向客户端推送消息
	for {
		select {
		case <-clientGone:
			ctx.Cancel(nil)
			return
		case message, ok := <-ch:
			if !ok || message == "complete" {
				return
			}
			// 发送消息
			c.SSEvent("message", message)
			c.Writer.Flush()
		}
	}
}

// CancelWorker 取消任务
func CancelWorker(c *gin.Context) {
	id := c.Query("id")
	if ctx, ok := workerContextMap[id]; ok {
		ctx.Cancel(nil)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
}
