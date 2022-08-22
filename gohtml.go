// @Time  : 2022/8/22 11:45
// @Email: jtyoui@qq.com

package gohtml

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml/core"
)

// NewFs create an Object from embed.FS file system.
//
// path == . indicates the current path.
func NewFs(efs embed.FS) gin.HandlerFunc {
	return core.NewFiler(newFs(efs), ".")
}

// New create an Object from Path file system.
func New(path string) gin.HandlerFunc {
	return core.NewFiler(newRead(path), path)
}
