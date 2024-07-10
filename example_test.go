package callerr_test

import (
	"fmt"

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

func Example() {
	if err := a(); err != nil {
		fmt.Println("Failed to a:", err)
	}
}
