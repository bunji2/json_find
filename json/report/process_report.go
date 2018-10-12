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

func Usage() string {
	return "Usage: %s report template report.json out.html\n"
}

func Process(args []string) (e error) {

	if len(args) < 3 {
		return errors.New("command report: too fiew arguments")
	}

	template_file := args[0]
	infile := args[1]
	outfile := args[2]

	var w *os.File
	w, e = os.Create(outfile)
	if e != nil {
		return
	}
	defer w.Close()

	var blob []byte
	if blob, e = ioutil.ReadFile(infile); e != nil {
		return
	}

	var json_tree reports_data
	if e = json.Unmarshal(blob, &json_tree); e != nil {
		return
	}

	e = process_report_template(w, template_file, json_tree)
	return
}

func process_report_template(w io.Writer, template_file string, data reports_data) (err error) {
	//fmt.Println("0")
	var tpl *template.Template
	//tpl, err = template.New("report_template").ParseFiles(template_file)
	tpl, err = template.ParseFiles(template_file)
	if err != nil {
		return
	}
	//fmt.Println("1")
	//err = tpl.ExecuteTemplate(w, "report_template", data)
	err = tpl.Execute(w, data)
	if err != nil {
		return
	}
	//fmt.Println("2")
	return
}

/*
func process_report(report report_data) (lines []string, e error) {
	lines = []string{}
	lines = append(lines, `<tr align="top">`)
	lines = append(lines,
		fmt.Sprintf(`<td align="left" rowspan="%d">%s</td>`,
			report.Summary.Total, report.Assume))
	lines = append(lines,
		fmt.Sprintf(`<td align="left" rowspan="%d">%s</td>`,
			report.Summary.Total, report.Note))
	lines = append(lines,
		fmt.Sprintf(`<td align="center" rowspan="%d">%d</td>`,
			report.Summary.Total, report.Summary.Total))
	lines = append(lines,
		fmt.Sprintf(`<td align="center" rowspan="%d">%d (%d)</td>`,
			report.Summary.Total, report.Summary.Ng, report.Summary.Rate))
	for _, line := range process_report_detail(report.Details[0]) {
		lines = append(lines, line)
	}
	lines = append(lines, `</tr>`)
	for _, detail := range report.Details[1:] {
		lines = append(lines, `<tr align="top">`)
		for _, line := range process_report_detail(detail) {
			lines = append(lines, line)
		}
		lines = append(lines, `</tr>`)
	}
	return
}

func process_report_detail(d detail_data) (lines []string) {
	lines = []string{}
	lines = append(lines,
		fmt.Sprintf(`<td align="center">%s -> %s</td>`, d.Route.From, d.Route.To))
	lines = append(lines,
		fmt.Sprintf(`<td align="center">%s</td>`, d.Communication.Result))
	lines = append(lines,
		fmt.Sprintf(`<td align="center">%s</td>`, d.Vulnerability.Result))
	return
}

func process_report_header() (lines []string) {
	return
}

func process_report_footer() (lines []string) {
	return
}
*/
