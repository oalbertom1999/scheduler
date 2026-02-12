package jobs

import (
	"fmt"
	"math/rand"
	"time"
)

// CleanupJob simulates a cleanup operation
func CleanupJob() {
	fmt.Printf("[%s] 🧹 Running cleanup job...\n", time.Now().Format("15:04:05"))
	
	// Simulate some work
	itemsCleaned := rand.Intn(100)
	time.Sleep(time.Millisecond * 500)
	
	fmt.Printf("[%s] ✅ Cleanup complete! Cleaned %d items\n", 
		time.Now().Format("15:04:05"), 
		itemsCleaned)
}
