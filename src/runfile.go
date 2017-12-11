package src

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
	Name     string                 `yaml:"name"`
	Details  string                 `yaml:"details"`
	Logging  string                 `yaml:"logging"`
	Backends []string               `yaml:"backends"`
	Plans    map[string]interface{} `yaml:"plan"`
}
