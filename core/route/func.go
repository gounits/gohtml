// @Time  : 2024/5/8 17:43
// @Email: jtyoui@qq.com

package route

// Func 声明一个自定义路由替换函数
// 将URL 替换成 resource
type Func func(url string) (resource string, err error)

func (f Func) Route(url string) (resource string, err error) {
	resource, err = f(url)
	return
}
