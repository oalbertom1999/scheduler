package jobs

import (
	"fmt"
	"time"
)

// HelloWorldJob is a simple job that prints a message
func HelloWorldJob() {
	fmt.Printf("[%s] 👋 Hello from scheduled job!\n", time.Now().Format("15:04:05"))
}
