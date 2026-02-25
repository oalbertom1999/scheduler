package jobs

import (
	"fmt"
	"time"
)

// HelloWorldJob prints a timestamped greeting to standard output.
// The timestamp is the local time formatted as "15:04:05" (HH:MM:SS).
func HelloWorldJob() {
	fmt.Printf("[%s] 👋 Hello from scheduled job!\n", time.Now().Format("15:04:05"))
}