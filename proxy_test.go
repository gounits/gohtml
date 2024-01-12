// Copyright 2023 Zhang Wei. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// @Time  : 2023/11/3 17:27
// @Email: jtyoui@qq.com

package gohtml_test

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml"
	"testing"
)

//go:embed html
var efs embed.FS

func TestNew(t *testing.T) {
	r := gin.Default()
	r.Use(gohtml.New(efs))
	// r.Use(gohtml.New("html"))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
