// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:37
// @Email: jtyoui@qq.com

package internal

import (
	"github.com/gounits/gohtml/core/file"
	"path"
	"path/filepath"
	"strings"
)

// realValidPath 获取到真实的路径
func realValidPath(dir string, address string) string {
	root := filepath.Join(dir, address)
	// 讲所有的 \\ 转为 / ,因为在 Fs中 \\ 是无法识别的
	root = strings.ReplaceAll(root, "\\", "/")
	return root
}

// FileResource 搜索资源路径，如果找到资源，则返回真实路径地址，
// 否则返回空字符串
// 搜索资源的算法：首先通过路由地址进行搜索
// 如果无法通过资源名称进行搜索
// 比如路由地址：/css/index.css 首先去css文件夹中搜索
// 如果无法搜索到css文件夹，则去其他文件夹中搜索
func FileResource(filer file.Filer, html string, dir string) string {
	if dirs, err := filer.ReadDir(dir); err == nil {
		// 获取真实有效的资源地址
		root := realValidPath(dir, html)

		// 判断资源地址是否存在
		if _, err = filer.Open(root); err == nil {
			return root
		}

		// 不存在递归搜索
		for _, address := range dirs {
			if address.IsDir() {
				// 这一步非常重要。 如果不进行拼接和验证有效路由，
				// 将会发生意外错误。
				newDir := realValidPath(dir, address.Name())
				return FileResource(filer, html, newDir)
			}
			// 获取资源名称
			name := path.Base(html)
			if address.Name() == name {
				return dir + "/" + name
			}
		}
	}
	return ""
}
