// squirrel allows for simple system configuration
// after an OS has booted for the first time. This is
// intended to run once and then self delete itself.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"github.com/michaeljs1990/squirrel/nuts"
	"github.com/michaeljs1990/squirrel/providers"
	"github.com/michaeljs1990/squirrel/src"
)

var (
	runfilePath = flag.String("runfile", "runfile.yaml", "path to a YAML runfile")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Infof("Attempting to load file from path: %s", *runfilePath)

	validateRunfilePath(*runfilePath)
	var b = readFile(*runfilePath)
	var runfile, err = src.NewRunfile(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	runfile.PrintHeader()
	providers.InitializeProviders(runfile.Backends)
	providers.FetchAllProviders()
	err = nuts.RunNuts(runfile.Plans)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Check if the path does exist and that it is a file
// or at least that we think it's a file.
func validateRunfilePath(path string) {
	fd, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File '%s' does not exist.\n", path)
			os.Exit(1)
		}
	}

	if fd.IsDir() {
		fmt.Printf("Path '%s' is not a file but a file is needed.\n", path)
		os.Exit(1)
	}

}

// Yes i'm being super lazy but these files
// shouldn't get that large ever and if they do
// i'll read this into a buffer first.
func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}
