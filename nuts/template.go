package nuts

import (
	"fmt"

	"github.com/golang/glog"
)

type Template struct {
	src  string
	dest string
	vars map[string]string
}

// ParseTemplate takes a chunck of YAML
// and turns it into a Template
// struct that is useful for later
func ParseTemplate(x map[interface{}]interface{}) (Module, error) {
	var t = Template{}

	for k, v := range x {
		var ok = true
		switch k {
		case "src":
			t.src, ok = v.(string)
		case "dest":
			t.dest, ok = v.(string)
		case "vars":
			var vars = make(map[string]string)
			var cast, ok = v.(map[interface{}]interface{})
			if ok {
				for ik, iv := range cast {
					vars[ik.(string)] = iv.(string)
				}
				t.vars = vars
			}
		default:
			glog.Infof("Template: %v", t)
			return t, fmt.Errorf("key '%s' is not valid in Template", k)
		}

		if !ok {
			glog.Infof("Template: %v", t)
			return t, fmt.Errorf("Trying to parse template module failed")
		}
	}

	return t, nil
}

// Run handles all logic fun the execution of
// this module
func (t Template) Run() error {
	return nil
}
