package jobs

import (
	"fmt"
	"time"
)

// TimeReportJob reports the current time every minute
func TimeReportJob() {
	now := time.Now()
	fmt.Printf("[%s] ⏰ Time Report - Current time: %s\n", 
		now.Format("15:04:05"), 
		now.Format("Monday, Jan 02, 2006 15:04:05 MST"))
}
