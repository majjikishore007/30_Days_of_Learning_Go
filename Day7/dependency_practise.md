# Go Dependency Injection Practice Exercises

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

## Exercise 2: Configuration Manager
**Goal**: Read configuration from different sources

```go
// Create a ConfigReader interface
type ConfigReader interface {
    ReadConfig() (map[string]string, error)
}

// Create a ConfigManager that accepts a ConfigReader
type ConfigManager struct {
    // Add your fields
}

// Implement different config readers:
// - FileConfigReader (reads from a file)
// - StringConfigReader (reads from a string, useful for testing)
// - EnvConfigReader (reads from environment variables)
```

**Test scenarios**:
- Test with StringConfigReader using hardcoded config
- Use FileConfigReader in main() to read from a JSON/YAML file
- Use EnvConfigReader to read from environment variables

## Exercise 3: Data Processor Pipeline
**Goal**: Process data through different transformations

```go
// Create interfaces for data processing
type DataReader interface {
    Read() ([]string, error)
}

type DataProcessor interface {
    Process(data []string) []string
}

type DataWriter interface {
    Write(data []string) error
}

// Create a Pipeline that accepts these dependencies
type Pipeline struct {
    // Add your fields
}

// Implement different processors:
// - UppercaseProcessor
// - FilterProcessor (filters out empty strings)
// - SortProcessor
```

**Test scenarios**:
- Test with mock/fake implementations
- Chain multiple processors together
- Read from strings.Reader, process, and write to different outputs

## Exercise 4: HTTP Client with Configurable Transport
**Goal**: Create an HTTP client wrapper with dependency injection

```go
// Create an HTTPClient interface
type HTTPClient interface {
    Get(url string) (*http.Response, error)
    Post(url string, body io.Reader) (*http.Response, error)
}

// Create a service that uses HTTPClient
type APIService struct {
    client HTTPClient
    baseURL string
}

// Implement methods like:
// - GetUser(id string) (*User, error)
// - CreateUser(user *User) error
```

**Test scenarios**:
- Create a mock HTTPClient for testing
- Use real http.Client in main()
- **Bonus**: Create a RetryHTTPClient that wraps another HTTPClient

## Exercise 5: Event System
**Goal**: Build a publish-subscribe system with injected dependencies

```go
// Create interfaces
type EventPublisher interface {
    Publish(event Event) error
}

type EventSubscriber interface {
    Subscribe(eventType string, handler func(Event)) error
}

// Create different implementations:
// - InMemoryEventBus
// - FileEventBus (logs events to file)
// - HTTPEventBus (sends events via HTTP)

type Event struct {
    Type      string
    Timestamp time.Time
    Data      interface{}
}
```

**Test scenarios**:
- Test with in-memory implementation
- Test event handling with multiple subscribers
- **Bonus**: Create a composite publisher that publishes to multiple destinations

## Exercise 6: Database Layer Abstraction
**Goal**: Create a data access layer with dependency injection

```go
// Create repository interfaces
type UserRepository interface {
    Create(user *User) error
    GetByID(id string) (*User, error)
    GetAll() ([]*User, error)
    Update(user *User) error
    Delete(id string) error
}

// Create a service that depends on the repository
type UserService struct {
    repo UserRepository
}

// Implement different repositories:
// - InMemoryUserRepository (for testing)
// - FileUserRepository (JSON file storage)
// - DatabaseUserRepository (if you want to use a real DB)
```

**Test scenarios**:
- Test UserService with InMemoryUserRepository
- Use FileUserRepository in main()
- Test error handling scenarios

## Exercise 7: Template Renderer
**Goal**: Create a flexible template rendering system

```go
// Create interfaces
type TemplateEngine interface {
    Render(templateName string, data interface{}) (string, error)
}

type TemplateLoader interface {
    LoadTemplate(name string) (string, error)
}

// Create a Renderer that combines these
type Renderer struct {
    engine TemplateEngine
    loader TemplateLoader
    writer io.Writer
}

// Implement different loaders:
// - FileTemplateLoader
// - StringTemplateLoader (for testing)
// - EmbeddedTemplateLoader
```

**Test scenarios**:
- Test with StringTemplateLoader and bytes.Buffer
- Render to files, HTTP responses, etc.
- **Bonus**: Add template caching

## Bonus Challenges

### Challenge 1: Middleware Pattern
Create a middleware system using dependency injection:

```go
type Handler interface {
    Handle(ctx context.Context, req *Request) (*Response, error)
}

type Middleware func(Handler) Handler

// Create middleware like:
// - LoggingMiddleware
// - AuthenticationMiddleware
// - RateLimitingMiddleware
```

### Challenge 2: Plugin System
Create a plugin architecture where plugins are injected:

```go
type Plugin interface {
    Name() string
    Execute(input interface{}) (interface{}, error)
}

type PluginManager struct {
    plugins []Plugin
}
```

## Tips for Success

1. **Start with interfaces**: Always define your interfaces first based on what behavior you need
2. **Test first**: Write your tests using fake/mock implementations before writing the real code
3. **Keep interfaces small**: Follow the Interface Segregation Principle - small, focused interfaces
4. **Use composition**: Build complex behavior by composing simple interfaces
5. **Think about the caller**: Design your interfaces from the perspective of the code that will use them

## Key Go Patterns to Practice

- **io.Writer/io.Reader**: For data flow
- **context.Context**: For cancellation and request-scoped values
- **Interface composition**: Combining multiple interfaces
- **Functional options**: For configurable constructors
- **Decorator pattern**: Wrapping implementations with additional behavior

Happy coding! Remember, the goal is to make your code testable, flexible, and reusable. Start with the simpler exercises and gradually work your way up to the more complex ones.