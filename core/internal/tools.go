// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:37
// @Email: jtyoui@qq.com

package internal

import "regexp"

// HasExtension 判断 URL 是否包含扩展名
func HasExtension(url string) bool {
	// 匹配扩展名（允许字母、数字，后面有点 .）
	re := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
	return re.MatchString(url)
}
