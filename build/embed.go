// +build ignore

// This script embeds `docopt.txt and `conf.yml` into the binary during at
// build time.

package main


import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {

	// get the cwd
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// get the project root
	root, err := filepath.Abs(cwd + "../../../")
	if err != nil {
		log.Fatal(err)
	}

	// specify template file information
	type file struct {
		In     string
		Out    string
		Method string
	}

	// enumerate the template files to process
	files := []file{
		file{
			In:     "cmd/cheat/docopt.txt",
			Out:    "cmd/cheat/str_usage.go",
			Method: "usage"},
		file{
			In:     "configs/conf.yml",
			Out:    "cmd/cheat/str_config.go",
			Method: "configs"},
	}

	// iterate over each static file
	for _, file := range files {

		// delete the outfile
		os.Remove(path.Join(root, file.Out))

		// read the static template
		bytes, err := ioutil.ReadFile(path.Join(root, file.In))
		if err != nil {
			log.Fatal(err)
		}

		// render the template
		data := template(file.Method, string(bytes))

		// write the file to the specified outpath
		spath := path.Join(root, file.Out)
		err = ioutil.WriteFile(spath, []byte(data), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// template packages the
func template(method string, body string) string {

	// specify the template string
	t := `package main

// Code generated .* DO NOT EDIT.

import (
	"strings"
)

func %s() string {
	return strings.TrimSpace(%s)
}
`

	return fmt.Sprintf(t, method, "`"+body+"`")
}
