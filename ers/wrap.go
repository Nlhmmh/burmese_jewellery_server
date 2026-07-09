package ers

import (
	"fmt"
	"runtime"
)

// wrap - Wraps an error with a stack trace for better debugging.
func wrap(err error) error {
	buf := make([]byte, 4096)
	return fmt.Errorf("%w\n%s", err, string(buf[:runtime.Stack(buf, false)]))
}
