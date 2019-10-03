package main

import "net/http"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type nodeType uint8
type Handle func(http.ResponseWriter, *http.Request, map[string]string)

const (
	static   nodeType = 0
	param    nodeType = 1
	catchAll nodeType = 2
)

type node struct {
	path      string
	wildChild bool
	nType     nodeType
	indices   []byte
	children  []*node
	handle    map[string]Handle
	priority  uint32
}

// func (n *node) addRoute(method, path, string, handle Handle) {
// 	n.priority++
// 	if len(n.path) > 0 || len(n.children) > 0 {
// 	WALK:
// 	}
// }

func (n *node) insertChild(method, path string, handle Handle) {
	var offset int

	for i, j := 0, len(path); i < j; i++ {
		if c := path[i]; c == ':' || c == '*' {
			if len(n.children) > 0 {
				panic("wildcard route conflicts with existing children")
			}

			k := i + 1
			for k < j && path[k] != '/' {
				k++
			}

			if k-i == 1 {
				panic("wildcards must be named with a non-empty name")
			}

			if c == ':' {
				if i > 0 {
					n.path = path[offset:i]
					offset = i
				}

				child := &node{
					nType: param,
				}
				n.children = []*node{child}
				n.wildChild = true
				n = child
				n.priority++

				if k < j {
					n.path = path[offset:k]
					offset = k
					child := &node{}
					n.children = []*node{child}
					n = child
					n.priority++
				}
			} else {
				if len(path) != k {
					panic("catch-all routes are only allowed at the end of the path")
				}

				if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
					panic("catch-all conflicts with existing handle for the path segment root")
				}

				i--
				if path[i] != '/' {
					panic("no / before catch-all")
				}

				n.path = path[offset:i]

				child := &node{
					wildChild: true,
					nType:     catchAll,
				}
				n.children = []*node{child}
				n.indices = []byte{path[i]}
				n = child
				n.priority++

				child = &node{
					path: path[i:],
					handle: map[string]Handle{
						method: handle,
					},
					nType:    catchAll,
					priority: 1,
				}
				n.children = []*node{child}

				return
			}
		}
	}

	n.path = path[offset:]
	n.handle = map[string]Handle{
		method: handle,
	}
}
