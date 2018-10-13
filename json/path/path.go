package path

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type patterns_data struct {
	Patterns []pattern_data `json:"patterns"`
}

type pattern_data struct {
	Assume string       `json:"assume"`
	Note   string       `json:"note"`
	Routes []route_data `json:"routes"`
}

type route_data struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var tpl *template.Template

func Init(template_file string) (e error) {
	tpl, e = template.ParseFiles(template_file)
	return
}

func Usage() string {
	return "Usage: %s path template file.json out.html\n"
}

func Process(args []string) (e error) {
	if len(args) < 3 {
		return errors.New("command path: too fiew arguments")
	}

	templ_file := args[0]
	infile := args[1]
	outfile := args[2]

	e = Init(templ_file)
	if e != nil {
		return
	}

	var w *os.File
	w, e = os.Create(outfile)
	if e != nil {
		return
	}
	defer w.Close()

	e = Write(w, infile)
	return
}

func Write(w io.Writer, infile string) (e error) {
	if tpl == nil {
		e = errors.New("command path: no template.")
		return
	}

	var blob []byte
	if blob, e = ioutil.ReadFile(infile); e != nil {
		return
	}

	var json_tree patterns_data
	if e = json.Unmarshal(blob, &json_tree); e != nil {
		return
	}

	e = tpl.Execute(w, json_tree)
	return
}
