package worker

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var CancelError = errors.New("cancel")

type Worker interface {
	Run()
}

type Context struct {
	ID       string
	Name     string
	Progress float64
	Logs     []Log
	Err      error
	Status   bool
}

const (
	// TypeInfo INFO
	TypeInfo = "info"
	// TypeWarn WARN
	TypeWarn = "warn"
	// TypeError ERROR
	TypeError = "error"
	// TypeData DATA
	TypeData = "data"
	// TypeProgress PROGRESS
	TypeProgress = "progress"
)

type Log struct {
	Time time.Time   `json:"time"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewContext(name string) *Context {
	return &Context{
		ID:       uuid.New().String(),
		Name:     name,
		Progress: 0,
		Logs:     make([]Log, 0),
		Err:      nil,
		Status:   true,
	}
}

func (ctx *Context) SetProgress(progress float64) {
	//保留四位小数
	progress = float64(int(progress*10000)) / 10000
	if progress > ctx.Progress {
		ctx.Progress = progress
		ctx.append(progress, TypeProgress)
	}
}

// Info 添加一条信息
func (ctx *Context) Info(text string) {
	ctx.append(text, TypeInfo)
}

// Warn 添加一条警告
func (ctx *Context) Warn(text string) {
	ctx.append(text, TypeWarn)
}

// Error 添加一条错误
func (ctx *Context) Error(text string) {
	ctx.append(text, TypeError)
}

func (ctx *Context) Data(data interface{}) {
	ctx.append(data, TypeData)
}

// log 添加一条日志
func (ctx *Context) append(data interface{}, typ string) {
	ctx.Logs = append(ctx.Logs, Log{
		Time: time.Now(),
		Type: typ,
		Data: data,
	})
}

func (ctx *Context) Complete() {
	ctx.Status = false
}

func (ctx *Context) Cancel(err error) {
	if err == nil {
		ctx.Err = CancelError
	} else {
		ctx.Err = err
	}
	ctx.Complete()
}

func (ctx *Context) IsCancel() bool {
	return !ctx.Status
}
