// @Time  : 2022/8/5 16:00
// @Author: jtyoui@outlook.com

package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml"
)

//go:embed fs
var efs embed.FS

func main() {
	r := gin.Default()
	r.Use(gohtml.NewFs(efs))
	//r.Use(gohtml.New("test/fs"))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
