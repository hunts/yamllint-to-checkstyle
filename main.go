package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	input    = os.Stdin
	output   = os.Stdout
	parsable = regexp.MustCompile(`^(?P<file>.+?):(?P<line>\d+):(?P<column>\d+):\s+\[(?P<level>[a-z]+)\]\s+(?P<desc>.+?)(?P<rule>\s+\([a-z-]+\))?\n$`)
)

type Checkstyle struct {
	XMLName xml.Name `xml:"checkstyle"`
	Version string   `xml:"version,attr"`
	Files   []*File
}

type File struct {
	XMLName  xml.Name `xml:"file"`
	Name     string   `xml:"name,attr"`
	Problems []Problem
}

type Problem struct {
	XMLName  xml.Name `xml:"error"`
	Line     int      `xml:"line,attr"`
	Column   int      `xml:"column,attr"`
	Severity string   `xml:"severity,attr"`
	Source   string   `xml:"source,attr,omitempty"`
	Message  string   `xml:"message,attr"`
}

func main() {
	reader := bufio.NewReader(input)
	err := process(reader, output)
	if err != nil {
		panic(err)
	}
}

func process(r *bufio.Reader, w io.Writer) error {
	fileMap := map[string]*File{}
	var fileOrder []string

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		match := parsable.FindStringSubmatch(line)
		if len(match) < 7 {
			return fmt.Errorf("unparsable line: %s", line)
		}

		filename, line_str, column_str, level, desc, rule := match[1], match[2], match[3], match[4], match[5], match[6]
		line_no, _ := strconv.Atoi(line_str)
		column_no, _ := strconv.Atoi(column_str)
		problem := Problem{Line: line_no, Column: column_no, Severity: level, Source: strings.Trim(rule, " ()"), Message: desc}

		file, ok := fileMap[filename]
		if !ok {
			fileOrder = append(fileOrder, filename)
			fileMap[filename] = &File{Name: filename, Problems: []Problem{problem}}
		} else {
			file.Problems = append(file.Problems, problem)
		}
	}

	var files []*File
	for _, filename := range fileOrder {
		files = append(files, fileMap[filename])
	}

	data, err := xml.MarshalIndent(Checkstyle{Version: "5.0", Files: files}, "", "    ")
	if err != nil {
		return err
	}

	fmt.Fprint(w, xml.Header)
	w.Write(data)
	w.Write([]byte("\n"))
	return nil
}
