package providers

// Provider must implement a basic set
// of data that ensures that you can expect
// your plans to run the same between them.
type Provider interface {
	Collect() map[string]interface{}
	GetCPUInfo()
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
