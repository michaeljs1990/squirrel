package providers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestLocalStruct is a simple test mostly useful for
// sanity when adding new features.
func TestLocalStruct(t *testing.T) {
	var l = NewLocal()

	l.Cores = 8
	l.Threads = 16
	l.VendorID = "Intel"
	l.Family = "Intel"
	l.Model = "Intel"
	l.ModelName = "Intel"

	var lmock = &Local{
		&CPU{
			Cores:     8,
			Threads:   16,
			VendorID:  "Intel",
			Family:    "Intel",
			Model:     "Intel",
			ModelName: "Intel",
		},
		&Host{},
	}

	if cmp.Equal(l, lmock) == false {
		t.Error("The struct l is not equal to lmock")
	}
}

// TestLocalCanCollectInfo ensure that we can collect
// CPU info without causing a panic in the go runtime.
// pass -v to go test to dump out the log which is a
// useful sanity check. Unfortunately the sys folder
// path for generating this output is hard coded. To
// Make this testable we need to upstream some changes
// to allow overwritting the sys path with some
// environment variables.
func TestLocalCanCollectInfo(t *testing.T) {
	var l = NewLocal()
	l.CollectCPUInfo()
	l.CollectHostInfo()
	t.Logf("Local Struct: %+v", l)
}

// TestLocalCanCollect is similar to CollectCPUs
// read above to see why we don't really even check
// anything inside this function.
func TestLocalCanCollect(t *testing.T) {
	var l = NewLocal()
	var c = l.Collect()
	t.Logf("Collection: %v", c)
}

func TestLocalIsAProvider(t *testing.T) {
	var l Provider = NewLocal()
	l.Collect()
}
