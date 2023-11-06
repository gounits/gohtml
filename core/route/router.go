// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 18:20
// @Email: jtyoui@qq.com

package route

/*
	路由规则接口

	Route 输入从http请求过来的URL地址
		   判断该地址是否满足路由规则，如果满足返回资源地址

根据路径获取到路由地址
*/
type Router interface {
	Route(url string) (string, error)
}
