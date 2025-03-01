package callerr_test

import (
	"github.com/jmank88/callerr"
)

func a() error {
	if err := b(); err != nil {
		return callerr.Format("failed to b: %w", err)
	}
	return nil
}
func b() error {
	if err := c(); err != nil {
		return callerr.Format("failed to c: %w", err)
	}
	return nil
}
func c() error {
	return callerr.New("original error")
}

func ExamplePrintln() {
	if err := a(); err != nil {
		callerr.Println("Failed to a: %v", err)
	}
}
