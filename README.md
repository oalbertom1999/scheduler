# Simple Cron Scheduler in Go

A simple yet powerful cron scheduler built with Go using the `robfig/cron` library.

## Features

- ✅ Schedule jobs using cron expressions
- ✅ Second-level precision
- ✅ Multiple example jobs included
- ✅ Graceful shutdown handling
- ✅ Clean and extensible architecture

## Installation

```bash
# Install dependencies
go mod download

# Run the scheduler
go run main.go

# Build the binary
go build -o scheduler
```

## Cron Expression Format

The scheduler uses the following format with second precision:

```
 ┌───────────── second (0 - 59)
 │ ┌───────────── minute (0 - 59)
 │ │ ┌───────────── hour (0 - 23)
 │ │ │ ┌───────────── day of month (1 - 31)
 │ │ │ │ ┌───────────── month (1 - 12)
 │ │ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
 │ │ │ │ │ │
 * * * * * *
```

### Examples:

- `*/10 * * * * *` - Every 10 seconds
- `0 * * * * *` - Every minute
- `0 */5 * * * *` - Every 5 minutes
- `0 0 9 * * *` - Every day at 9:00 AM
- `0 0 10 * * MON` - Every Monday at 10:00 AM
- `0 30 8 * * MON-FRI` - Every weekday at 8:30 AM

## Project Structure

```
scheduler/
├── main.go           # Entry point and job registration
├── jobs/             # Job implementations
│   ├── hello_world.go
│   ├── time_report.go
│   ├── cleanup.go
│   ├── daily_report.go
│   └── email_job.go
├── go.mod
└── README.md
```

## Included Jobs

1. **HelloWorldJob** - Runs every 10 seconds, prints a simple greeting
2. **TimeReportJob** - Runs every minute, displays current time
3. **CleanupJob** - Runs every 5 minutes, simulates cleanup operations
4. **DailyReportJob** - Runs every day at 9:00 AM, generates a report
5. **EmailJob** - Runs every Monday at 10:00 AM, demonstrates struct-based jobs

## Creating Custom Jobs

### Simple Function Job

```go
func MyCustomJob() {
    fmt.Println("My custom job is running!")
}

// Register in main.go
c.AddFunc("0 */15 * * * *", jobs.MyCustomJob)
```

### Struct-based Job (Advanced)

```go
type MyJob struct {
    Config string
}

func (j *MyJob) Run() {
    fmt.Printf("Running with config: %s\n", j.Config)
}

// Register in main.go
myJob := &jobs.MyJob{Config: "production"}
c.AddJob("0 0 * * * *", myJob)
```

## Usage

```bash
# Run the scheduler
go run main.go

# Output:
# 🚀 Cron scheduler started!
# 📋 Scheduled jobs:
#    - Next run: 2026-02-12 16:54:30
#    - Next run: 2026-02-12 16:55:00
# ...

# Press Ctrl+C to stop gracefully
```

## Dependencies

- [robfig/cron](https://github.com/robfig/cron) - Cron library for Go

## License

MIT
