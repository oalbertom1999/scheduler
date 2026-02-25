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

// main is the program entry point. It creates a cron scheduler with second-level precision, registers example jobs, starts the scheduler and prints upcoming run times, then blocks until SIGINT or SIGTERM and performs a graceful shutdown.
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

// registerJobs registers a set of example scheduled jobs on the provided cron scheduler.
// 
// It schedules:
// - HelloWorldJob to run every 10 seconds.
// - TimeReportJob to run at the start of every minute.
// - CleanupJob to run every 5 minutes.
// - DailyReportJob to run daily at 09:00.
// - EmailJob (struct) to run every Monday at 10:00 with a preset recipient and subject.
//
// Any error returned by the scheduler when adding a job is logged fatally, causing the program to exit.
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

	// Example 6: Run every 30 seconds
	_, err = c.AddFunc("*/30 * * * * *", jobs.ReportProcessorJob)
	if err != nil {
		log.Fatalf("Error scheduling ReportProcessorJob: %v", err)
	}

	// Example 7: Run every 30 seconds
	_, err = c.AddFunc("*/30 * * * * *", jobs.ProcessUserDataJob)
	if err != nil {
		log.Fatalf("Error scheduling ProcessUserDataJob: %v", err)
	}
}