package providers

import "fmt"

// Provider must implement a basic set
// of data that ensures that you can expect
// your plans to run the same between them.
type Provider interface {
	Collect() map[string]interface{}
}

// Package level vars so we don't have to make expensive
// network or system calls to get these multiple times.
var local map[string]interface{}

// InitializeProviders looks at the list of
// providers you asked for in your YAML and
// attempts to set them all for you in a
// global scope for when you ask for them in
// the future.
func InitializeProviders(s []string) {
	local = NewLocal().Collect()
}

// FetchAllProviders merges all providers that
// have been set and returns them to you for use.
func FetchAllProviders() map[string]interface{} {
	for k, v := range local {
		fmt.Printf("%s : %v\n", k, v)
	}
	return local
}
