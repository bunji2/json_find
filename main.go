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

	var data interface{}
	if err = json.Unmarshal(blob, &data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	//	fmt.Println("data#0.Name = ", data.([]interface{})[0].(map[string]interface{})["Name"].(string))

	//fmt.Println(read_path1("#0.Name"))
	//	fmt.Println(json_find(data, "#0.Name"))
	value, ok := json_find(data, pattern)
	if !ok {
		fmt.Fprintln(os.Stderr, "(not found)")
		return 2
	}
	fmt.Println(value)
	return 0
}

/*
func run() int {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0])
		return 1
	}

	filepath := os.Args[1]

	var err error
	var blob []byte
	if blob, err = ioutil.ReadFile(filepath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	var data interface{}
	if err = json.Unmarshal(blob, &data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	for _, obj := range data.([]interface{}) {
		o := obj.(map[string]interface{})
		for k, v := range o {
			v, _ = v.(string)
			fmt.Println(k, v)
		}
	}

	fmt.Println("data#0.Name = ", data.([]interface{})[0].(map[string]interface{})["Name"].(string))

	//fmt.Println(read_path1("#0.Name"))
	fmt.Println(json_find(data, "#0.Name"))
	fmt.Println(json_find(data, "#1.Order"))
	return 0
}
*/
