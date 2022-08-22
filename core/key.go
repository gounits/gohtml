// @Time  : 2022/8/22 13:36
// @Email: jtyoui@qq.com

package core

import "github.com/gin-gonic/gin"

/*
	Filer load file interface

	core
		1. ReadDir Read the file according to the path, return whether the directory under the file is a file and its name
		2. Open Read the contents of the file according to the path

Both fs.go and path.go implement this interface
*/
type Filer interface {
	ReadDir(dir string) ([]Stat, error)
	Open(path string) ([]byte, error)
}

type Stat struct {
	dir  bool
	name string
}

func NewStat(isDir bool, name string) Stat {
	return Stat{dir: isDir, name: name}
}

func (s Stat) IsDir() bool {
	return s.dir
}

func (s Stat) Name() string {
	return s.name
}

// NewFiler The terminal where it all started
func NewFiler(f Filer, path string) gin.HandlerFunc {
	return NewDist(path).LoadFs(f)
}
