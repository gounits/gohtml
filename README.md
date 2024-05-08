# goHTML

    一种静态代理技术，结合Gin框架一起使用，快速高效。

## 1. install
     go get github.com/gounits/gohtml

## 2. usage

![img.png](img/img.png)

```go
package main


import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gounits/gohtml"
	"github.com/gounits/gohtml/core/route"
	"testing"
)


//go:embed html
var efs embed.FS

func TestNew(t *testing.T) {
	r := gin.Default()
	r.Use(gohtml.NewProxy(efs, route.Default))
	// r.Use(gohtml.New("html"))
	// r.Use(gohtml.New(efs))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}



```

![img_2.png](img/img_2.png)
## 3.Web
    click http://localhost:8080/

![img_1.png](img/img_1.png)
