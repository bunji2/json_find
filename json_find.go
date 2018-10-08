package main

import (
	"strconv"
	"strings"
)

// json ツリーから特定のノードを取り出す。
// #N --- data[N]
// .X --- data.X
// Example
// data = [
//    {"Name": "Platypus", "Order": "Monotremata"},
//    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]
// json_find(data, "#0.Name") ==> ("Platypus", true)
// json_find(data, "#1.Order") ==> ("Dasyuromorphia", true)

func json_find(data interface{}, path string) (r interface{}, ok bool) {
	p := split_path(path)
	//fmt.Println("path =", p)
	n := data
	for _, x := range p {
		//fmt.Println("x =", x, "x[0] =", string(x[0]))
		switch x[0] {
		case '#':
			i, e := strconv.Atoi(x[1:])
			if e == nil {
				//fmt.Printf("#%d\n", i)
				q, ok2 := elem(n, i)
				if !ok2 {
					return
				}
				n = q
			}
		case '.':
			name := x[1:]
			//fmt.Printf(".%s\n", name)
			q, ok2 := prop(n, name)
			if !ok2 {
				return
			}
			n = q
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
	/*
		switch n.(type) {
		case []interface{}:
			ok = true
			r = n.([]interface{})[i]
			return
		}
	*/
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
	/*
		switch n.(type) {
		case map[string]interface{}:
			ok = true
			r, ok = n.(map[string]interface{})[name]
			return
		}
	*/
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
