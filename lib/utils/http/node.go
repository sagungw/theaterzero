package http

import (
	"golang.org/x/net/html"
)

// DeleteNode :nodoc:
func DeleteNode(node *html.Node, remove *html.Node) (err error) {
	for n := node.FirstChild; n != nil; n = n.NextSibling {

	}

	return
}
