// @Time  : 2022/8/22 11:38
// @Email: jtyoui@qq.com

package gohtml

import (
	"github.com/gounits/gohtml/core"
	"os"
)

type read struct {
	path string
}

func newRead(path string) core.Filer {
	return &read{path}
}

func (f *read) ReadDir(dir string) ([]core.Stat, error) {
	var stat []core.Stat

	if dirs, err := os.ReadDir(dir); err != nil {
		return nil, err
	} else {
		for _, d := range dirs {
			stat = append(stat, core.NewStat(d.IsDir(), d.Name()))
		}
	}

	return stat, nil
}

func (f *read) Open(path string) ([]byte, error) {
	return os.ReadFile(path)
}
