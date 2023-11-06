// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:03
// @Email: jtyoui@qq.com

package file

import (
	"embed"
	"errors"
)

type fs struct {
	efs embed.FS
}

func NewFs(efd embed.FS) Filer {
	return &fs{efd}
}

func (f *fs) ReadDir(dir string) ([]Stater, error) {
	var ef []Stater
	values, err := f.efs.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, value := range values {
		ef = append(ef, value)
	}

	return ef, nil
}

func (f *fs) Open(path string) ([]byte, error) {
	data, err := f.efs.Open(path)
	if err != nil {
		return nil, err
	}

	stats, err := data.Stat()
	if err != nil {
		return nil, err
	}

	if stats.IsDir() {
		return nil, errors.New("is Dir")
	}

	text, err := f.efs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if _, err = data.Read(text); err != nil {
		return nil, err
	}

	return text, nil
}
