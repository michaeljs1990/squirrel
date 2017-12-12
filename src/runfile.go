package src

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Runfile specifies what a valid yaml
// configuration looks like for input
// into a squirrel program. Sorry I can't
// figure out how to break this up better so
// the parsing for each plan will get passed
// to the module and they have to figure out
// what to do.... Not great but I don't see
// anyway around this in go. The mess has to live
// in the Runfile struct or the Parser. It's
// easier to isolate it in the Parser IMO.
type Runfile struct {
	Name      string                   `yaml:"name"`
	Details   string                   `yaml:"details"`
	Backends  []string                 `yaml:"backends"`
	Codenames []string                 `yaml:"codenames"`
	Plans     []map[string]interface{} `yaml:"plan"`
}

// NewRunfile takes a slice of bytes and trys
// to make it into a Runfile struct
func NewRunfile(b []byte) (Runfile, error) {
	var r = Runfile{}
	var err = yaml.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// PrintHeader prints out info about your
// runfile at the start of a run.
func (r Runfile) PrintHeader() {
	fmt.Println()
	fmt.Println("==============================================================")
	fmt.Println("Name: " + r.Name)
	fmt.Println("Details:  " + r.Details)
	fmt.Println("Backends: ")
	for _, v := range r.Backends {
		fmt.Println(" - " + v)
	}
	fmt.Println("Codenames: ")
	for _, v := range r.Codenames {
		fmt.Println(" - " + v)
	}
	fmt.Println("==============================================================")
}
