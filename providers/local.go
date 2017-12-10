// Package providers will fetch information about the current
// system that you are running on for use inside commands.
package providers

import (
	"github.com/fatih/structs"
	"github.com/shirou/gopsutil/cpu"
)

// Local ...
type Local struct {
	*CPU
}

// NewLocal ...
func NewLocal() *Local {
	return &Local{
		&CPU{},
	}
}

// Collect all the information about the Local struct
// here and convert it to a map to pass back. This is
// so we can easily provide this as input to a template
// later
func (l *Local) Collect() map[string]interface{} {
	l.GetCPUInfo()
	var ret = structs.Map(l)
	return ret
}

// GetCPUInfo fetches all known information
// about the CPU and packages it up for your
// convenience
func (l *Local) GetCPUInfo() {
	var cpus, err = cpu.Info()
	if err != nil {
		panic(err)
	}

	var cpu = cpus[0]
	l.Cores = cpu.Cores
	l.Threads = cpu.Cores * 2
	l.VendorID = cpu.VendorID
	l.Family = cpu.Family
	l.Model = cpu.Model
	l.ModelName = cpu.ModelName
}
