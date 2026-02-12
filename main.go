package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oalbertom99/scheduler/jobs"
	"github.com/robfig/cron/v3"
)

func main() {
	// Create a new cron scheduler with second precision
	c := cron.New(cron.WithSeconds())

	// Register jobs
	registerJobs(c)

	// Start the scheduler
	c.Start()
	fmt.Println("🚀 Cron scheduler started!")
	fmt.Println("📋 Scheduled jobs:")
	for _, entry := range c.Entries() {
		fmt.Printf("   - Next run: %v\n", entry.Next)
	}

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutting down scheduler...")
	c.Stop()
	fmt.Println("✅ Scheduler stopped gracefully")
}

func registerJobs(c *cron.Cron) {
	// Example 1: Run every 10 seconds
	_, err := c.AddFunc("*/10 * * * * *", jobs.HelloWorldJob)
	if err != nil {
		log.Fatalf("Error scheduling HelloWorldJob: %v", err)
	}

	// Example 2: Run every minute
	_, err = c.AddFunc("0 * * * * *", jobs.TimeReportJob)
	if err != nil {
		log.Fatalf("Error scheduling TimeReportJob: %v", err)
	}

	// Example 3: Run every 5 minutes
	_, err = c.AddFunc("0 */5 * * * *", jobs.CleanupJob)
	if err != nil {
		log.Fatalf("Error scheduling CleanupJob: %v", err)
	}

	// Example 4: Run every day at 9:00 AM
	_, err = c.AddFunc("0 0 9 * * *", jobs.DailyReportJob)
	if err != nil {
		log.Fatalf("Error scheduling DailyReportJob: %v", err)
	}

	// Example 5: Using a struct job (more advanced)
	emailJob := &jobs.EmailJob{
		Recipients: []string{"admin@example.com"},
		Subject:    "Weekly Report",
	}
	_, err = c.AddJob("0 0 10 * * MON", emailJob) // Every Monday at 10:00 AM
	if err != nil {
		log.Fatalf("Error scheduling EmailJob: %v", err)
	}
}
