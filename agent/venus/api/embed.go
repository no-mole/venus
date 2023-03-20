package api

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"io"
	"net/http"
	"path"
	"strings"
	"sync"
)

//go:embed dist/*
var f embed.FS

var index string
var loadIndexError error

var readIndexOnce = &sync.Once{}

func loadIndex() error {
	indexFile, err := f.Open("index.html")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(indexFile)
	if err != nil {
		return err
	}
	index = string(data)
	return nil
}

func UIHandle(ctx *gin.Context) {
	readIndexOnce.Do(func() {
		loadIndexError = loadIndex()
		if loadIndexError != nil {
			loadIndexError = fmt.Errorf("load index file err:%s", loadIndexError)
		}
	})
	if loadIndexError != nil {
		output.Json(ctx, loadIndexError, nil)
		return
	}
	filePath := path.Clean(strings.TrimLeft(ctx.Request.URL.Path, "/"))
	//因为路由问题，不存在的文件请求路径全部返回index.html
	if _, err := f.Open(filePath); err != nil {
		ctx.Header("content-type", "text/html;charset=utf-8")
		ctx.String(200, index)
		return
	}
	http.FileServer(http.FS(f)).ServeHTTP(ctx.Writer, ctx.Request)
}
