// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 15:07
// @Email: jtyoui@qq.com

package file

/*
	文件状态接口

	IsDir 判断文件是否是文档类型
	Name 获取文件的文件名称

用于返回 Filer 接口
*/
type Stater interface {
	IsDir() bool
	Name() string
}

/*
	加载文件需要实现的接口


	ReadDir 根据输入dir的文件路径读取文件夹，然后返回该文件夹下面所有文件状态

	Open 根据输入的path路径来读取文件

读取静态资源需要实现的核心接口
*/
type Filer interface {
	ReadDir(dir string) ([]Stater, error)
	Open(path string) ([]byte, error)
}
