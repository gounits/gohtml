// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:42
// @Email: jtyoui@qq.com

package rule

type Ruler interface {
	Match(url string, path string) bool
}
