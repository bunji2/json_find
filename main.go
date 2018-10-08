package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	Usage = `Usage: %s json_file pattern+
pattern := # <number of index>
        |  . <property name>
`
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0])
		return 1
	}

	filepath := os.Args[1]
	pattern := os.Args[2]

	var err error
	var blob []byte
	if blob, err = ioutil.ReadFile(filepath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	var json_tree interface{}
	if err = json.Unmarshal(blob, &json_tree); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	//	fmt.Println("json_tree#0.Name = ", json_tree.([]interface{})[0].(map[string]interface{})["Name"].(string))

	//fmt.Println(read_path1("#0.Name"))
	//	fmt.Println(json_find(json_tree, "#0.Name"))
	value, ok := json_find(json_tree, pattern)
	if !ok {
		fmt.Fprintln(os.Stderr, "(not found)")
		return 2
	}
	fmt.Println(value)
	return 0
}
