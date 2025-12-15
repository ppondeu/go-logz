# logz

`logz` is a lightweight Go logging package built on top of [Uber Zap](https://github.com/uber-go/zap). It provides **structured logging**, **formatted logs**, and supports **lazy initialization** and **re-initialization** for flexible usage in Go projects.  

---

## **Features**

- Automatic **default logger** if `InitLog` is not called.  
- Supports **re-initialization** to update logger configuration at runtime.  
- Easy-to-use structured logging with `Info`, `Debug`, `Warn`, `Error`, `Fatal`.  
- Supports formatted logging via `Infof`, `Errorf`, `Fatalf`.  
- Configurable for **development** and **production** environments.  
- Thread-safe and safe for concurrent use.  

---

## **Installation**

```bash
go get github.com/ppondeu/go-logz@v0.1.2
```

## **Usage**

### 1. Basic Usage (auto-initialized default logger)
You can call log functions without initializing explicitly:
```go
package main

import "github.com/ppondeu/go-logz"

func main() {
	logz.Info("Default logger active")
	logz.Debug("Debugging info")
	logz.Warn("This is a warning")
	logz.Errorf("Error message with format: %v", "example")
}
```
```The default logger is in development mode.```

### 2. Custom Initialization
Initialize the logger explicitly for production or custom configuration:
```go
logz.InitLog("production")
logz.Info("Using production logger")
```

### 3. Re-initialization with Options
You can re-initialize the logger at any time, for example to change the log level or output path:
```go
logz.InitLog("development", func(c *zap.Config) {
	c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	c.OutputPaths = []string{"stdout", "./app.log"}
})
logz.Debug("Now using debug logger with custom output")
```
serverMode can be "development" or "production".
Optional functions allow customization of zap.Config.

### 4. Logging Functions
#### Structured Logging
```go
logz.Info("User created", zapcore.Field{Key: "user_id", Type: zapcore.StringType, String: "1234"})
logz.Debug("Debug info")
logz.Warn("Warning message")
logz.Error("Something went wrong")
logz.Fatal("Critical error, exiting")
```

#### Formatted Logging
```go
logz.Infof("User %s has logged in", username)
logz.Errorf("Failed to process request: %v", err)
logz.Fatalf("Fatal error: %s", errMsg)
```

### Example
```go
package main

import (
	"github.com/ppondeu/go-logz"
	"go.uber.org/zap"
)

func main() {
	// Lazy default logger
	logz.Info("Starting application")

	// Custom initialization
	logz.InitLog("production")
	logz.Info("Production logger active")

	// Re-initialize with debug level and custom output
	logz.InitLog("development", func(c *zap.Config) {
		c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		c.OutputPaths = []string{"stdout", "./app.log"}
	})
	logz.Debug("Debug log with custom output")
}
```