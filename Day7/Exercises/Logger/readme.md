## Exercise 1: Logger with Different Outputs
**Goal**: Create a logger that can write to different destinations

```go
// Create a Logger struct that accepts an io.Writer
type Logger struct {
    // Add your fields here
}

// Implement methods:
// - Info(message string)
// - Error(message string)
// - Debug(message string)

// Each method should prefix the message with timestamp and log level
// Example: "2023-07-03 14:30:00 [INFO] Your message here"
```

**Test scenarios**:
- Log to a bytes.Buffer for testing
- Log to a file in main()
- Log to stdout in main()
- **Bonus**: Log to multiple destinations simultaneously using `io.MultiWriter`