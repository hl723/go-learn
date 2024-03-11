# Go Coding Conventions

This is my notes from [effective go](https://go.dev/doc/effective_go) documentation.


## Formatting
- Use `gofmt` or `go fmt`!

- Tabs for indentation
- Long Lines - wrap and indent once
- C-Style block and line comments (`/* */ and //`)

## Names
Uppercase names are exported while lowercase remains private within a package.

Use MixedCaps/CamelCase for multiword names

### Package Names
- Import package names is the accessor for the contents
- Should be short, concise, single word
- Are the basename of a path `encoding/base64` would be referenced as `base64` in code

### Getters/Setters
- Getters - Just name the getter receiver.Fieldname 
- Setters - receiver.SetFieldname

### Interfaces
- One method interfaces should have -er suffixes. Examples: Reader, Writer, Formatter, CloseNotifier...
- String not ToString

### Semicolons
- Are automatically inserted.
- If, for, switch, select statements need to have left brace on the same line!

## Control Structures
Can take an initializer variable.

- Function params and return values belong to function scope.

### For
- for init; condition; pos {}
- for index range slice
- for index, value range slice
- for key range map
- for key, value range map
- for rune in string  // NOT BYTE

### Switch
- Use switch statements for long if-then-else statements
- Cases can be in comma-separated lists. 
```
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```
### Type Switches
Used to dynamically type an interface variable.
```
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## Functions

### Multiple Return Values
Syntax `func (file *File) Write(b []byte) (n int, err error)`

- The result params can be named. If the function is short, can just call return, otherwise, need to call return with the names for readability

### Defer
Defer a function call until end of function. E.g. for mutexes and closing files.


## Data
It is OK to return an address of a local variable!

### new
Allocation of zero initialized memory can be done with `new(x)`. It returns a pointer.

- `new(File) and &File{}` are equivalent

### make
Used only for slices, maps, channels.
Make does *not* return a pointer.
- make([]type, length, (optional)capacity)
- new([]type) creates a slice tuple (nil, type) where the underlying array is nil.
- make(map[keyType]valueType)
- make(chan dataType[, bufsize])

### Arrays
- Arrays are values, passing an array to a function passes a copy.
- Size of array is part of its type. The types [10]int and [20]int are distinct.

### Maps
- Test presence with `_, present := map[key]`
- Use `delete(map, key)` to delete

## Init Functions
Can have one or more init functions to set up state.
- Process order: Init import packages -> Variable declarations and their inits -> Init functions.

## Blank Identifiers
- Can be used to silence unused imports/variables during development
- Import for side effect `import _ "net/http/pprof"`

## Concurrency
Go is made for concurrency, not necessarily parallelization.

### Goroutines
Lightweight execution of a function in a shared memory space. 
It is multiplexed onto multiple OS threads.

### Channels
Can be used as semaphores. Producers dump data into the channel, Consumers read from the channel. The channel buffer size is the number of active calls.

## Panic (throw an exception)
Stop the execution of the current function due to an unrecoverable error.

Then unwinds the goroutine stack, runs any deferred functions along the way. If rewinding goes to top of call stack, program dies.

`panic(errorMsg)`

### Recover (catch)
Control can be regained if a deferred function calls `recover`.

So `recover` is only useful in deferred functions.
