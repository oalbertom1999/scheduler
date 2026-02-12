package jobs

import (
	"fmt"
	"strings"
	"time"
)

// EmailJob is a more complex job that implements the cron.Job interface
type EmailJob struct {
	Recipients []string
	Subject    string
}

// Run implements the cron.Job interface
func (e *EmailJob) Run() {
	fmt.Printf("[%s] 📧 Sending email job...\n", time.Now().Format("15:04:05"))
	fmt.Printf("   To: %s\n", strings.Join(e.Recipients, ", "))
	fmt.Printf("   Subject: %s\n", e.Subject)
	
	// Simulate sending email
	time.Sleep(time.Millisecond * 300)
	
	fmt.Printf("[%s] ✅ Email sent successfully!\n", time.Now().Format("15:04:05"))
}
