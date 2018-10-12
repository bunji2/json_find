package indent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func Usage() string {
	return "Usage: %s indent file.json\n"
}

func Process(args []string) (e error) {
	if len(args) < 1 {
		return errors.New("command indent: too fiew arguments")
	}

	infile := args[0]

	var blob []byte
	if blob, e = ioutil.ReadFile(infile); e != nil {
		return
	}

	var json_tree interface{}
	if e = json.Unmarshal(blob, &json_tree); e != nil {
		return
	}

	if blob, e = json.MarshalIndent(json_tree, "", "  "); e != nil {
		return
	}

	fmt.Print(string(blob))
	return
}
