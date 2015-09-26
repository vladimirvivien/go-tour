package main

import "fmt"

type node struct {
	data string
}

type tree map[node]interface{}

func search(key node, t tree) interface{} {
	for k, v := range t {
		if k.data == key.data {
			return v
		}
		switch val := v.(type) {
		case node:
			if val.data == key.data {
				return val
			}
		case map[node]interface{}:
			return search(key, val)
		}
	}
	return node{}
}

func (t tree) graph(key node, value interface{}) {
	found := search(key, t)
	// insert new
	if (found == node{}) {
		t[key] = value
		return
	}

	// update existing val
	switch val := found.(type) {
	case map[node]interface{}:
		subTree, ok := value.(map[node]interface{})
		if !ok {
			panic("Value must be of type map[node]interface{}")
		}
		for k, v := range subTree {
			val[k] = v
		}

	default:
		fmt.Println("Value not found, inserting", key, value)
		val = value
	}

}

func main() {
	t := map[node]interface{}{
		node{"A"}: map[node]interface{}{
			node{"B"}: "Insane!",
			node{"C"}: node{"12"},
			node{"D"}: map[node]interface{}{
				node{"E"}: 44,
				node{"F"}: node{"G"},
			},
		},
	}
	tree(t).graph(node{"AA"}, "Simplee")
	tree(t).graph(node{"D"}, map[node]interface{}{node{"H"}: 100, node{"I"}: map[node]interface{}{node{"K"}: 90}})
	tree(t).graph(node{"AA"}, map[node]interface{}{node{"S"}: "Simplified"})
	//fmt.Println(search(node{"D"}, t))
	fmt.Println(t)
}
