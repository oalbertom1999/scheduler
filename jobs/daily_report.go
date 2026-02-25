package jobs

import (
	"fmt"
	"math/rand"
	"time"
)

// DailyReportJob generates and prints a simulated daily report to standard output with a timestamped header.
// The report includes Total Users (500–1499), Active Users (0 up to Total Users-1), and Revenue (0 ≤ revenue < 10000), and uses the current time for the displayed timestamp and date.
func DailyReportJob() {
	fmt.Printf("[%s] 📊 Generating daily report...\n", time.Now().Format("15:04:05"))
	
	// Simulate report generation
	totalUsers := rand.Intn(1000) + 500
	activeUsers := rand.Intn(totalUsers)
	revenue := rand.Float64() * 10000
	
	fmt.Println("========================")
	fmt.Printf("📅 Daily Report - %s\n", time.Now().Format("Jan 02, 2006"))
	fmt.Println("========================")
	fmt.Printf("Total Users: %d\n", totalUsers)
	fmt.Printf("Active Users: %d\n", activeUsers)
	fmt.Printf("Revenue: $%.2f\n", revenue)
	fmt.Println("========================")
}