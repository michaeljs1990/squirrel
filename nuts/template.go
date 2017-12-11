package nuts

import "errors"

type Template struct {
	src  string
	dest string
	vars map[string]interface{}
}

// ParseTemplate takes a chunck of YAML
// and turns it into a Template
// struct that is useful for later
func ParseTemplate(x map[string]interface{}) (Module, error) {
	var t = Template{}
	return t, errors.New("lol")
}

// Run handles all logic fun the execution of
// this module
func (t Template) Run() error {
	return nil
}
