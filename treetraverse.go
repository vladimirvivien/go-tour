package main

import "fmt"

type node struct {
	data  string
	nodes []*node
}

func (n *node) String() string {
	return fmt.Sprintf("%v -> %v", n.data, n.nodes)
}

type tree []*node

// search for one match and stops
func search(key *node, root tree) *node {
	// width search at root level first
	for _, n := range root {
		if key.data == n.data {
			return n
		}
	}

	// depth search
	for _, n := range root {
		if n.nodes != nil {
			return search(key, n.nodes)
		}
	}

	return nil
}

func graph(t tree, n *node) tree {
	found := search(n, t)
	if found != nil {
		found = n
		return t
	}
	t = append(t, n)
	return t
}

func main() {
	var t tree = []*node{
		&node{
			"A",
			[]*node{
				&node{data: "B"},
				&node{data: "C"},
				&node{"D", []*node{&node{"E", nil}}},
			},
		},
	}

	fmt.Println("Search for node{D, nil}", search(&node{"D", nil}, t))
	t = graph(t, &node{"F", nil})
	fmt.Println("Search for node{F, nil}", search(&node{"F", nil}, t))
}
