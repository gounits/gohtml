// @Time  : 2022/8/22 10:37
// @Email: jtyoui@qq.com

package gohtml

import (
	"embed"
	"errors"
	"github.com/gounits/gohtml/core"
)

type fs struct {
	efs embed.FS
}

func newFs(efd embed.FS) *fs {
	return &fs{efd}
}

func (f *fs) ReadDir(dir string) ([]core.Stat, error) {
	var ef []core.Stat
	values, err := f.efs.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, value := range values {
		ef = append(ef, core.NewStat(value.IsDir(), value.Name()))
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
	if _, err = data.Read(text); err != nil {
		return nil, err
	}

	return text, nil
}
