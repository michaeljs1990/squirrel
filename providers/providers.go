package providers

// Provider must implement a basic set
// of data that ensures that you can expect
// your plans to run the same between them.
type Provider interface {
	Collect() map[string]interface{}
	CollectCPUInfo()
	CollectHostInfo()
}

// CPU basic information
type CPU struct {
	Cores     int32
	Threads   int32
	VendorID  string
	Family    string
	Model     string
	ModelName string
}

// Host returns basic information
// about the OS you are running on
type Host struct {
	OS                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	VirtualizationSystem string
	VirtualizationRole   string
}

// Disk is not yet implemented because I can't
// find any good way to get a list of block devices
// in golang along with the sizes. I'll add the
// functionality for this is a different repo since
// it will likely be a non trivial amount of code
// just for linux support.
type Disk struct{}
