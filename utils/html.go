package utils

import "golang.org/x/net/html"

func FindTagByName(node *html.Node, name string) (result *html.Node) {
	if node.Data == name {
		return node
	}
	for currChild := node.FirstChild; result == nil && currChild != nil; currChild = currChild.NextSibling {
		result = FindTagByName(currChild, name)
	}
	return result
}
