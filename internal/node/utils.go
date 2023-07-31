package node

import "strings"

func Path(node Node, path string) []Node {
	prev := []Node{node}

	sp := strings.Split(path, ".")
	for _, key := range sp {
		var current []Node
		for _, prevNode := range prev {
			if key == "*" {
				list := prevNode.List()
				current = append(current, list...)
			} else {
				get, err := prevNode.Get(key)
				if err != nil {
					continue
				}
				current = append(current, get)
			}
		}
		prev = current
	}
	return prev
}
