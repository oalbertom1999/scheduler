package jobs

import (
	"fmt"
	"math/rand"
	"time"
)

// CleanupJob simulates a cleanup operation by logging a start message, waiting briefly, and logging a completion message with the number of items cleaned.
// It writes two lines to standard output with timestamps formatted as "15:04:05", sleeps for 500 milliseconds, and reports a pseudo-random cleaned count between 0 and 99.
func CleanupJob() {
	fmt.Printf("[%s] 🧹 Running cleanup job...\n", time.Now().Format("15:04:05"))
	
	// Simulate some work
	itemsCleaned := rand.Intn(100)
	time.Sleep(time.Millisecond * 500)
	
	fmt.Printf("[%s] ✅ Cleanup complete! Cleaned %d items\n", 
		time.Now().Format("15:04:05"), 
		itemsCleaned)
}