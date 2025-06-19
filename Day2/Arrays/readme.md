# Key points

declaring an array in go 

fixed_size 
[N]type{values, values, ...}

dynamic sizes 
[...]type{values, values, ...} // for a slice

comparing slieces dont work with ==
we need to use reflect.DeepEqual(slice1, slice2) cannot gurantee typpe safety
- len funciton 
- make function 
- variadic functions (dynamic arguments sizes)
example 
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```