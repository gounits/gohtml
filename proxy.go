// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:27
// @Email: jtyoui@qq.com

package gohtml

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml/core"
	"github.com/gounits/gohtml/core/file"
)

// NewHTML 创建 Gin HandlerFunc
func NewHTML(filer file.Filer, auto *core.AutoReplace) gin.HandlerFunc {
	return auto.Load(filer)
}

// New 加载本地文件
//
// obj 如果是一个字符串 那么就加载本地路径
// 如果是 embed.FS 类型 那么就加载 Fs 文件
func New(obj interface{}) gin.HandlerFunc {
	switch path := obj.(type) {
	case string:
		dist := core.NewDist(path)
		return NewHTML(file.NewLocal(), dist)
	case embed.FS:
		dist := core.NewDist(".")
		return NewHTML(file.NewFs(path), dist)
	default:
		panic("不支持该类型")
	}
	return nil
}
