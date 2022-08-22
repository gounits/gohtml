// @Time  : 2022/8/5 15:35
// @Author: jtyoui@outlook.com

package core

// Rule Rules table from url
type Rule struct {
	Router   string     // routing address
	Resource string     // static resource address in dist
	Model    MatchModel // replace pattern
}

// DefaultRule The default rule, routing starts from the root directory
func defaultRule() *Rule {
	r := &Rule{
		Router:   "/",
		Resource: "index.html",
		Model:    AccurateMatching,
	}
	return r
}

// checkRouter check if the route is in resource.
//
// Resource is HTML.Dist in data path.
// url is network access address.
func (r *Rule) checkRouter(url string) bool {
	return r.Model.match(url, r.Router)
}
