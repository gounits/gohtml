// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:37
// @Email: jtyoui@qq.com

package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml/core/file"
	"github.com/gounits/gohtml/core/internal"
	"github.com/gounits/gohtml/core/route"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"
)

type AutoReplace struct {
	rules    []route.Router // 路由规则表
	rootPath string         // 资源的根目录
}

// NewDist 初始化规则表结构
func NewDist(path string) *AutoReplace {
	return &AutoReplace{rootPath: path}
}

// AddRules 增加路由获取接口
func (a *AutoReplace) AddRules(r ...route.Router) {
	a.rules = append(a.rules, r...)
}

// CustomRuleFunc 添加自定义替换函数
// 首先采用自定义替换函数。
func (a *AutoReplace) CustomRuleFunc(f route.Func) {
	a.rules = append(a.rules, f)
}

func (a *AutoReplace) AddRuleFunc(r route.Func) {
	a.rules = append(a.rules, r)
}

/*
	muxPath 根据不同的参数获取不同的静态网页地址

	路径 { "/" : "index.html" , "/home" : "home.html" }
	例如： / --> index.html /home --> home.html

如果路由无法获取自定义资源名称，则返回路由地址
*/
func (a *AutoReplace) muxPath(url string) (resource string, err error) {
	for _, rule := range a.rules {
		resource, err = rule.Route(url)
		if resource != "" && err == nil {
			return
		}
	}
	return
}

// static 多叉路由解析静态文件代理
func (a *AutoReplace) static(filer file.Filer) gin.HandlerFunc {
	agent := func(c *gin.Context) {
		var (
			html       string // 获取的网页地址URL
			staticFile string // 获取真实的文件信息地址
			err        error  // 异常
			data       []byte // 找到的资源数据
		)

		// 先走API接口
		c.Next()

		// 当获取API接口为404的是否，才考虑走代理
		if c.Writer.Status() != http.StatusNotFound {
			return
		}

		// 获取上一页的路由地址
		url := strings.TrimSpace(c.Request.URL.Path)

		// 根据路径映射表，从路由地址中获取静态网页名称
		if html, err = a.muxPath(url); err != nil {
			if errors.Is(err, route.NotMatchError) {
				// 如果没有找到对应的路由地址，则直接返回原始URL地址
				html = url
			}
		}

		// 判断静态资源是否存在，存在则返回真实路径，不存在返回空字符串
		if staticFile = internal.FileResource(filer, html, a.rootPath); staticFile == "" {
			// 如果地址包含静态数据，则直接返回。
			// 判断地址是否请求.css .js .png 等类似文件 是就直接返回
			if internal.HasExtension(url) || internal.HasExtension(html) {
				return
			}

			// 没有找到对应的资源，则返回根路由
			if staticFile, err = a.muxPath("/"); err != nil {
				return
			}
		}

		// 如果 staticFile 是 Dir，则停止
		// 注意：映射的路由表不能有目录名字
		if _, err = filer.ReadDir(staticFile); err == nil {
			return
		}

		// 再次读取资源，如果无法读取则立即结束
		if data, err = filer.Open(staticFile); err != nil {
			log.Printf("【ERROR】没有读取到资源： %v", err)
			return
		}

		// 根据静态资源或路由名称获取后缀名称
		// 例如：index.html -> html , index.js -> js
		suffix := path.Ext(html)

		// 根据后缀返回不同的Content-Type响应
		c.Header("Content-Type", mime.TypeByExtension(suffix))

		// 写回文件的内容
		c.Status(http.StatusOK)

		if _, err = c.Writer.Write(data); err != nil {
			log.Printf("【ERROR】返回响应数据失败： %v", err)
			return
		}
		c.Abort()
	}
	return agent
}

// Load 加载静态文件
func (a *AutoReplace) Load(filer file.Filer) gin.HandlerFunc {
	return a.static(filer)
}
