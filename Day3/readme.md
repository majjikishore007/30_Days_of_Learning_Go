# Topic Structs, methods & interfaces


## Structs 
- Struct is a named collection of feilds that can hold data of different types.

Declaring a struct
```go
type Circle struct {
    Radius float64
}
```

## Methods and Functions

### method
- method is a function with a receiver
- syntax `func (receiverName ReceiverType) MethodName(args)`

### function
- function is a standalone block of code that can be called independently
- syntax `func FunctionName(args)`

- Method overlading is not supported in Go,

- When your method is called on a variable of that type, you get your reference to its data via the receiverName variable. In many other programming languages this is done implicitly and you access the receiver via this.

### Choice we have
- Declare a different package and put same method in it
- Define methods with different names


## Interfaces
- Interface is a type that defines a set of methods signatures without providing the actual implementation.
- It defines a contract that needs to be fulfilled by any type that implements the interface.
- In Go they are highly important because they allow for polymorphism, enabling different types and create highly-decoupled code with type-sfety.
- In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.


## Anonymous Structs
- Anonymous structs are structs that are defined without a name.
```go
areaTest := []struct {
    name string
    radius float64
}{
    {"Circle", 5.0},
    {"Circle", 10.0},
}
```

# Wrapping up

- Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
- Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
- Adding methods so you can add functionality to your data types and so you can implement interfaces
- Table based tests to make your assertions clearer and your suites easier to extend & maintain
