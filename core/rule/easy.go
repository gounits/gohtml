// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 18:08
// @Email: jtyoui@qq.com

package rule

import (
	"regexp"
	"strings"
)

// Easy 匹配规则
type Easy uint8

const (
	FuzzyMatching    Easy = iota // 模糊匹配
	AccurateMatching             // 精准匹配
	RegexpMatching               // 规则匹配
	SuffixMatching               // 后缀匹配
	PrefixMatching               // 前缀匹配
	AbsoluteMatching             // 绝对匹配
)

// Match URL地址和路由根据规则去判断，是否成功
func (e Easy) Match(url string, router string) (ok bool) {
	switch e {
	case FuzzyMatching:
		ok = strings.Contains(url, router)
	case AccurateMatching:
		ok = url == router
	case RegexpMatching:
		ok = regexp.MustCompile(router).MatchString(url)
	case SuffixMatching:
		ok = strings.HasSuffix(url, router)
	case PrefixMatching:
		ok = strings.HasPrefix(url, router)
	case AbsoluteMatching:
		ok = true
	default:
		ok = false
	}
	return
}
