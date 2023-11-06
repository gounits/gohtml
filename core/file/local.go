// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:19
// @Email: jtyoui@qq.com

package file

import "os"

type local struct {
}

func NewLocal() Filer {
	return new(local)
}

func (*local) ReadDir(dir string) ([]Stater, error) {
	var stat []Stater
	dirs, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		stat = append(stat, d)
	}

	return stat, nil
}

func (*local) Open(path string) ([]byte, error) {
	return os.ReadFile(path)
}
