// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 18:29
// @Email: jtyoui@qq.com

package route

import (
	"errors"
	"github.com/gounits/gohtml/core/rule"
)

// ruleRoute 定义一个简单的规则路由
// 路由地址 一一对应 一个静态资源的地址
// Router -> Resource
type ruleRoute struct {
	Router   string     // 路由的地址
	Resource string     // 静态资源的地址
	Rules    rule.Ruler // 规则匹配是否满足
}

var (
	NotMatchError = errors.New("路由不匹配")
)

func New(router string, resource string, rule_ rule.Ruler) Router {
	return &ruleRoute{
		Router:   router,
		Resource: resource,
		Rules:    rule_,
	}
}

func (r *ruleRoute) Route(url string) (path string, err error) {
	if r.Rules.Match(url, r.Router) {
		path = r.Resource
		return
	}
	err = NotMatchError
	return
}

// Default 定义一个默认的规则路由，此路由用处比较多
// 在很多的Js框架，都会编译成 dist/index.html
var Default = New("/", "index.html", rule.AccurateMatching)
