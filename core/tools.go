// @Time  : 2022/8/22 9:50
// @Email: jtyoui@qq.com

package core

import (
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type MatchModel uint8

const (
	FuzzyMatching    MatchModel = iota // fuzzy matching
	AccurateMatching                   // exact match
	RegexpMatching                     // regular match
)

// Different rules based on replacement pattern
func (m MatchModel) match(url, router string) (ok bool) {
	switch m {
	case FuzzyMatching:
		ok = strings.Contains(url, router)
	case AccurateMatching:
		ok = url == router
	case RegexpMatching:
		ok = regexp.MustCompile(router).MatchString(url)
	default:
		ok = false
	}
	return
}

// head Responses to different Content-Types according to different types of resources
func head(suffix string) string {
	typ := ""
	switch strings.ToLower(suffix) {
	case ".html", ".htm", ".css":
		typ = "text/" + suffix[1:]
	case ".js":
		typ = "application/javascript"
	case ".ico":
		typ = "image/x-icon"
	case ".png", ".jpg", ".jpeg":
		typ = "image/" + suffix[1:]
	case ".woff", ".woff2":
		typ = "font/" + suffix[1:]
	default:
		typ = "application/json; charset=UTF-8"
	}
	return typ
}

// realValidPath get real path
func realValidPath(dir string, address string) string {
	root := filepath.Join(dir, address)
	// All path symbols need to be converted from \ to /
	root = strings.ReplaceAll(root, "\\", "/")
	return root
}

// fileResource Search the resource path, if the resource is found, return the real path address,
// otherwise return an empty string
// Algorithm for searching resources: first search by the address of the route,
// if you can't search by the name of the resource
// For example, routing address: /css/index.css First go to the css folder to search,
// if the css folder cannot be searched, go to other folders to search
func fileResource(filer Filer, html string, dir string) string {
	if dirs, err := filer.ReadDir(dir); err == nil {
		// Get the real and effective resource address
		root := realValidPath(dir, html)

		// Determine whether the resource address exists
		if _, err = filer.Open(root); err == nil {
			return root
		}

		// does not exist to search recursively
		for _, address := range dirs {
			if address.IsDir() {
				// This step is very important. If you do not splice and verifying valid routes,
				// unexpected errors will occur.
				newDir := realValidPath(dir, address.Name())
				return fileResource(filer, html, newDir)
			} else {
				// Get the name of the resource
				name := path.Base(html)
				if address.Name() == name {
					return dir + "/" + name
				}
			}
		}
	}
	return ""
}
