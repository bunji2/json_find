package report

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type reports_data struct {
	Reports []report_data `json:"reports"`
}

type report_data struct {
	Assume  string        `json:"assume"`
	Note    string        `json:"note"`
	Summary summary_data  `json:"summary"`
	Details []detail_data `json:"detail"`
}

type summary_data struct {
	Total int `json:"total"`
	Ng    int `json:"ng"`
	Rate  int `json:"rate"`
}

type detail_data struct {
	Route         route_data  `json:"route"`
	Communication result_data `json:"communication"`
	Vulnerability result_data `json:"vulnerability_result"`
}

type route_data struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type result_data struct {
	Result string `json:"result"`
	Report string `json:"report"`
}

var tpl *template.Template

func Init(template_file string) (e error) {
	tpl, e = template.ParseFiles(template_file)
	return
}

func Usage() string {
	return "Usage: %s report template report.json out.html\n"
}

func Process(args []string) (e error) {

	if len(args) < 3 {
		e = errors.New("command report: too fiew arguments")
		return
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
		e = errors.New("command report: no template.")
		return
	}

	var blob []byte
	if blob, e = ioutil.ReadFile(infile); e != nil {
		return
	}

	var json_tree reports_data
	if e = json.Unmarshal(blob, &json_tree); e != nil {
		return
	}

	e = tpl.Execute(w, json_tree)
	return
}
