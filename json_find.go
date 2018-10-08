package main

import (
	"strconv"
	"strings"
)

// This function finds the node that matches pattern from json tree.
// json ツリーから pattern に該当するノードを見つける関数
// params:
//   json_tree --- json tree
//   pattern --- see following.
// pattern := pat+
// pat := #N | .X
//   #N --- json_tree[N]
//   .X --- json_tree.X
// Example:
// json_tree = [
//    {"Name": "Platypus", "Order": "Monotremata"},
//    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]
// json_find(json_tree, "#0.Name") ==> ("Platypus", true)
// json_find(json_tree, "#1.Order") ==> ("Dasyuromorphia", true)
// json_find(json_tree, "#2.Name") ==> ("", false)

func json_find(json_tree interface{}, path string) (r interface{}, ok bool) {
	p := split_path(path)
	//fmt.Println("path =", p)
	n := json_tree
	for _, x := range p {
		//fmt.Println("x =", x, "x[0] =", string(x[0]))
		switch x[0] {
		case '#':
			i, e := strconv.Atoi(x[1:])
			if e != nil {
				return
			}
			//fmt.Printf("#%d\n", i)
			r, ok = elem(n, i)
			if !ok {
				return
			}
			ok = false
			n = r
		case '.':
			name := x[1:]
			//fmt.Printf(".%s\n", name)
			r, ok = prop(n, name)
			if !ok {
				return
			}
			ok = false
			n = r
		default:
			return
		}
	}
	r = n
	ok = true
	return
}

func elem(n interface{}, i int) (r interface{}, ok bool) {
	var t []interface{}
	t, ok = n.([]interface{})
	if !ok {
		return
	}
	if i < 0 || i >= len(t) {
		ok = false
		return
	}
	r = t[i]
	return
}

func prop(n interface{}, name string) (r interface{}, ok bool) {
	var t map[string]interface{}
	t, ok = n.(map[string]interface{})
	if !ok {
		return
	}
	ok = false
	r, ok = t[name]
	return
}

func split_path(path string) (r []string) {
	a := strings.Split(path, "#")
	if a[0] == "" {
		a = a[1:]
	}
	r = []string{}
	for _, x := range a {
		for _, y := range split_path2(x) {

			if y[0] == '.' {
				r = append(r, y)
			} else {
				r = append(r, "#"+y)
			}
		}
	}
	return
}

func split_path2(path string) []string {
	t := strings.Split(path, ".")
	r := []string{}
	if t[0] != "" {
		r = append(r, t[0])
	}

	for _, x := range t[1:] {
		r = append(r, "."+x)
	}
	return r
}
