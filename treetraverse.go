package main

import "fmt"

type node struct {
	data  string
	nodes []*node
}

type tree []*node


func search(key *node, root tree) *node {
	fmt.Println("Searching tree:", root)
	for _, n := range root {
		if key.data == n.data {
			return n
		}
		if n.nodes != nil {
			return search(key, n.nodes)
		}
	}

	return nil
}

func (t tree) graph(n *node) {
	fmt.Println("Graphing", n)
	found := search(n, t)
	if found != nil {
		fmt.Println("Graphed node exists, update")
		found = n
		return
	}
	fmt.Println("Graph node not found, adding")
	fmt.Println("tree before graph", t)
	var k []*node = append(t, n)
	t = k
	fmt.Println("tree after graph", t)
}

func main() {
var	t tree = []*node{
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
	fmt.Println("Search for node{F, nil}", search(&node{"F", nil}, t))
	t.graph(&node{"F", nil})
	fmt.Println(t)
	fmt.Println("Search for node{F, nil}", search(&node{"F", nil}, t))
}
