// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:37
// @Email: jtyoui@qq.com

package internal

import "strings"

// Header 根据不同类型的资源响应不同的Content-Type
func Header(suffix string) string {
	typ := ""
	switch strings.ToLower(suffix) {
	case ".html", ".htm", ".css":
		typ = "text/" + suffix[1:]
	case ".js":
		typ = "application/javascript"
	case ".ico":
		typ = "image/x-icon"
	case ".png", ".jpg", ".jpeg":
		typ = "image/" + suffix[1:]
	case ".woff", ".woff2":
		typ = "font/" + suffix[1:]
	default:
		typ = "application/json; charset=UTF-8"
	}
	return typ
}
