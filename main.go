package main

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

var source = flag.String("src", path.Join(".", "in.csv"), "Input data file path.")
var target = flag.String("tar", path.Join(".", "out.txt"), "Out data file path.")
var tplF = flag.String("tpl", path.Join(".", "out.tpl"), "Go template file.")
var skipHead = flag.Bool("skip", true, "Skip header line.")

func main() {
	flag.Parse()

	// init template
	f, err := ioutil.ReadFile(*tplF)
	if err != nil {
		logExit(fmt.Sprintf("error read template file: %v", err), 1)
	}

	tpl, err := template.New("gage").Funcs(template.FuncMap{
		"GUID": GUID,
	}).Parse(string(f))

	if err != nil {
		logExit(fmt.Sprintf("error parse template file: %v", err), 1)
	}

	// read data
	in, err := ioutil.ReadFile(*source)
	if err != nil {
		logExit(fmt.Sprintf("error read input data file: %v", err), 1)
	}

	r := csv.NewReader(bytes.NewReader(in))
	lines, err := r.ReadAll()
	if err != nil {
		logExit(fmt.Sprintf("error read csv file: %v", err), 1)
	}

	rows := make([]map[string]string, 0)
	for i, line := range lines {
		if i == 0 && *skipHead {
			continue
		}
		data := make(map[string]string)
		for j, col := range line {
			data[fmt.Sprintf("col%d", j)] = col
		}
		rows = append(rows, data)
	}

	buf := bytes.NewBuffer([]byte{})
	if err := tpl.Execute(buf, rows); err != nil {
		logExit(fmt.Sprintf("error execute tpl: %v", err), 1)
	}

	// print out
	out, err := os.Create(*target)
	if err != nil {
		logExit(fmt.Sprintf("error create out file : %v", err), 1)
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := buf.WriteTo(out); err != nil {
		logExit(fmt.Sprintf("error wtrite to out file : %v", err), 1)
	}
}

// GUID - randomly generated hexadecimal string of the form ac1b73df-5e35-f216-b6ef-4957b9bc5781
func GUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err == nil {
		return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	}
	return ""
}

func logExit(msg string, code int) {
	fmt.Printf("%s\n", msg)
	os.Exit(code)
}
