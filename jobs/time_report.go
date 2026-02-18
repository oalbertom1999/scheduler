package jobs

import (
	"fmt"
	"time"
)

// TimeReportJob prints a single timestamped line to standard output containing the current time.
// The output is prefixed with the current 24-hour time in brackets and includes the full weekday, date,
// time, and timezone (e.g., "[15:04:05] ⏰ Time Report - Current time: Monday, Jan 02, 2006 15:04:05 MST").
func TimeReportJob() {
	now := time.Now()
	fmt.Printf("[%s] ⏰ Time Report - Current time: %s\n", 
		now.Format("15:04:05"), 
		now.Format("Monday, Jan 02, 2006 15:04:05 MST"))
}