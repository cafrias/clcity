package elapsed

import (
	"fmt"
	"time"
)

// Elapsed returns function that when executed counts the time in seconds from the time when Elapsed was executed.
func Elapsed() func() {
	start := time.Now()
	return func() {
		ts := fmt.Sprintf("%.2fs", time.Since(start).Seconds())
		fmt.Printf("--------------------------------\nFinished in %s\n", ts)
	}
}
