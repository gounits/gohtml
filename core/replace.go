// @Time  : 2022/8/22 9:45
// @Email: jtyoui@qq.com

package core

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"strings"
)

type AutoReplace struct {
	rules    []*Rule
	f        func(url string) string
	findPath string // find the beginning of a path Rule.Resource
}

// NewDist Initialize the rule table structure
//
// LoadFs Mount the content in go fs, the request url will be intercepted first, and then look for the resources in fs
func NewDist(path string) *AutoReplace {
	return &AutoReplace{findPath: path}
}

// AddRules Add rule table to replace resources
func (a *AutoReplace) AddRules(rule ...*Rule) {
	a.rules = append(a.rules, rule...)
}

// CustomRuleFunc Add custom replacement function
// The custom replacement function is taken first.
// If the custom function returns an empty string,
// then the rule table will be taken (if the rule table is defined)
func (a *AutoReplace) CustomRuleFunc(f func(url string) (resource string)) {
	a.f = f
}

/*
	muxPath Get different static web page addresses according to different parameters

	paths { "/" : "index.html" , "/home" : "home.html" }
	eg： / --> index.html  /home --> home.html

If the route cannot get the custom resource name, return the route address
*/
func (a *AutoReplace) muxPath(url string) string {
	if a.f != nil {
		resource := a.f(url)
		if resource != "" {
			return resource
		}
	}

	for _, rule := range a.rules {
		if rule.checkRouter(url) {
			return rule.Resource
		}
	}
	return url
}

// static Multipart static files
func (a *AutoReplace) static(filer Filer) gin.HandlerFunc {
	agent := func(c *gin.Context) {
		// Get the routing address of the previous page
		url := strings.TrimSpace(c.Request.URL.Path)

		// According to the paths mapping table, get the static web page name from the routing address
		// If there is no direct return route
		html := a.muxPath(url)

		// Determine whether the static resource exists, and return the real path if it exists
		// does not exist returns an empty string
		staticFile := fileResource(filer, html, a.findPath)
		if staticFile == "" {
			return
		}

		// Continue if staticFile is Dir
		if _, err := filer.ReadDir(staticFile); err == nil {
			return
		}

		// Read the resource again, if it cannot be read, end immediately
		text, err := filer.Open(staticFile)
		if err != nil {
			log.Println("Static resources are found, but resource information cannot be obtained: " + err.Error())
			return
		}

		// Get the suffix name based on the static resource or route name
		// eg: index.html -> html , index.js -> js
		suffix := path.Ext(html)

		// Block different responses Content-Type based on the suffix name
		c.Header("Content-Type", head(suffix))

		// write back the contents of the file
		c.Status(http.StatusOK)
		_, err = c.Writer.Write(text)
		if err != nil {
			log.Println("Failed to send static resources: " + err.Error())
			return
		}
		c.Abort()
	}
	return agent
}

// LoadFs Default index.html static file
func (a *AutoReplace) LoadFs(filer Filer) gin.HandlerFunc {
	if a.f == nil && a.rules == nil {
		a.AddRules(defaultRule())
	}
	return a.static(filer)
}
