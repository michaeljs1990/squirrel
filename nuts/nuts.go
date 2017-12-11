package nuts

import (
	"fmt"
	"os"

	"github.com/michaeljs1990/squirrel/src"
)

// Nut provides the interface that must
// be implemented for adding new features
// to the Plan struct.
type Nut interface {
	Parse(map[interface{}]interface{}) (Module, error)
}

// Module is the interface needed so that we
// can run the Nut that has been parsed.
type Module interface {
	Run() error
}

// RunNuts glues all the nuts together with
// the runfile. This is pretty ugly but I guess
// it's not super magical which is cool...
func RunNuts(x src.Runfile) error {
	var modules []Module
	for _, v := range x.Plans {
		for ik, iv := range v {
			// I'm taking advantate of switch having no fallthrough
			// in go so I don't have to check for err != nil inside
			// every single case statement
			var vm, ok = iv.(map[interface{}]interface{})
			if !ok {
				fmt.Printf("Can't convert %v to map[string]interface.\n", v)
				os.Exit(1)
			}

			var err error
			var m Module

			switch ik {
			case "template":
				m, err = ParseTemplate(vm)
				modules = append(modules, m)
			default:
				return fmt.Errorf("'%s' did not match any current modules", ik)
			}

			if err != nil {
				return err
			}

		}
	}

	// Run all the parsed modules
	for _, m := range modules {
		m.Run()
	}

	return nil
}
