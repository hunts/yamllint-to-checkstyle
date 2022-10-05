package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	inDir  = "testdata/input/"
	outDir = "testdata/output/"
)

func TestProcess(t *testing.T) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		t.Errorf("failed to read the testdata: %s", err)
	}

	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			inpath := inDir + file.Name()
			outpath := outDir + file.Name()

			f, err := os.Open(inpath)
			if err != nil {
				t.Errorf("failed to load input: %s", err)
			}

			in := bufio.NewReader(f)
			out := &bytes.Buffer{}

			process(in, out)

			expectedBytes, err := ioutil.ReadFile(outpath)
			if err != nil {
				t.Errorf("failed to load output: %s", err)
			}

			actual := out.String()
			expected := string(expectedBytes)
			if strings.Compare(expected, actual) != 0 {
				t.Errorf("expected:\n\n%s\nbut got:\n\n%s\n", expected, actual)
			}
		})
	}
}
