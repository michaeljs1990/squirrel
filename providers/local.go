// Package providers will fetch information about the current
// system that you are running on for use inside commands.
package providers

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

// Local ...
type Local struct {
	vars map[string]interface{}
}

// NewLocal ...
func NewLocal() *Local {
	var v = make(map[string]interface{})
	return &Local{vars: v}
}

// Collect all the information about the Local struct
// here and convert it to a map to pass back. This is
// so we can easily provide this as input to a template
// later
func (l *Local) Collect() map[string]interface{} {
	l.CollectCPUInfo()
	l.CollectHostInfo()

	var localMap = make(map[string]interface{})
	for k, v := range l.vars {
		localMap["local_"+k] = v
	}
	return localMap
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
	l.vars["cores"] = cpu.Cores
	l.vars["threads"] = cpu.Cores * 2
	l.vars["vendorID"] = cpu.VendorID
	l.vars["family"] = cpu.Family
	l.vars["models"] = cpu.Model
	l.vars["model_name"] = cpu.ModelName
}

// CollectHostInfo sets all info regarding
// OS and Platform
func (l *Local) CollectHostInfo() {
	var host, err = host.Info()
	if err != nil {
		panic(err)
	}

	l.vars["os"] = host.OS
	l.vars["platform"] = host.Platform
	l.vars["platform_family"] = host.PlatformFamily
	l.vars["platform_version"] = host.PlatformVersion
	l.vars["virtualization_system"] = host.VirtualizationSystem
	l.vars["virtualization_role"] = host.VirtualizationRole
}
