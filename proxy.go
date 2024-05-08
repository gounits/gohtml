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
	"github.com/gounits/gohtml/core/route"
)

// New 加载本地文件
//
// obj 如果是一个字符串 那么就加载本地路径
// 如果是 embed.FS 类型 那么就加载 Fs 文件
func New(obj interface{}) gin.HandlerFunc {
	return NewProxy(obj, route.Default)
}

func NewProxy(obj interface{}, rules ...route.Router) gin.HandlerFunc {
	var (
		filer   file.Filer
		replace *core.AutoReplace
	)

	switch path := obj.(type) {
	case string:
		replace = core.NewDist(path)
		filer = file.NewLocal()
	case embed.FS:
		replace = core.NewDist(".")
		filer = file.NewFs(path)
	default:
		panic("不支持该类型")
	}

	// 替换规则
	replace.AddRules(rules...)
	return replace.Load(filer)
}
