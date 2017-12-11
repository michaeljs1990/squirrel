// Package providers will fetch information about the current
// system that you are running on for use inside commands.
package providers

import (
	"github.com/fatih/structs"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

// Local ...
type Local struct {
	*CPU
	*Host
}

// NewLocal ...
func NewLocal() *Local {
	return &Local{
		&CPU{},
		&Host{},
	}
}

// Collect all the information about the Local struct
// here and convert it to a map to pass back. This is
// so we can easily provide this as input to a template
// later
func (l *Local) Collect() map[string]interface{} {
	l.CollectCPUInfo()
	l.CollectHostInfo()

	var ret = structs.Map(l)
	return ret
}

// CollectCPUInfo fetches all known information
// about the CPU and packages it up for your
// convenience
func (l *Local) CollectCPUInfo() {
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

// CollectHostInfo sets all info regarding
// OS and Platform
func (l *Local) CollectHostInfo() {
	var host, err = host.Info()
	if err != nil {
		panic(err)
	}

	l.OS = host.OS
	l.Platform = host.Platform
	l.PlatformFamily = host.PlatformFamily
	l.PlatformVersion = host.PlatformVersion
	l.VirtualizationSystem = host.VirtualizationSystem
	l.VirtualizationRole = host.VirtualizationRole
}
