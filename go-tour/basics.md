# A Tour of Go - Basics

Go over a summary of the basics of Go.

## Packages 
- Each program starts in `package main`

- Your program can use other packages with `import "path/to/package"`
- The package name is the same as the last element in the import path. `math/rand` uses files from that begin with `package rand`
- Names are exported if it begins with a captial letter.

## Functions 
-  `func name(param1 type1, ...) returnType`

- If params share types, can omit from all, but the last. Like `func f(x, y, z int)`
- Can return multiple values by having a tuple of return types `(int, string)`
- Return values can be named. Those are treated like variables. A return statement without args will return the named return values implicitly. This is known as a "naked" return. Should only be used in short functions!

## Variables
- Declare a variable with `var name type`

- If initializer is present, can omit type, type will be inferred. `var a, b, c = 1,true, "hi"`
- `:=` Can declare and init a variable only *inside* a function!
- Uninitialized values are given 0 for numerics, false for boolean, "" for strings
- Type conversions can be done with `T(v)` where T is the dst type
- Constants can be declared with `const PI = 3.14`, cannot use := syntax
- Untyped constants takes the type needed by its context

## Variable Types
- bool 

- string
- int, uint, uintptr - defaults to 32-bits on 32 bit systems, 64 bit on 64-bit systems
- int8, int16, int32, int64, and also uint versions
- byte - alias for uint8
- rune - alias for int32, represents a Unicode "letter"
- float32, float64
- complex64, complex128 - complex numbers (a + bi)

## Loops and Conditionals
- Only for loops, no while. C-for loop.

- For a while loop, just write for. `for sum < 1000`
- Infinite Loop `for {}`
- if statements can run a short statement before the condition separated by semicolon. That variable is only within the scope of the if and else block. `if v := math.Pow(x, n); v < lim`
- Switch statements are similar to C. Except no break statements are needed. Evaluated from top to bottom.
- Switch can be run without condition. A substitute for long if-then-else statements.
- Defer statement defers execution of the statement until the surrounding function returns. The arguments are evaluated immediately.
- Defer function calls are placed in a stack and executed in LIFO manner.

## Pointers
- Like C, there are * and & operations

- There is no pointer arithmetic.

## Structs
```
type Vertex struct {
    X int 
    Y int
}

// Struct creation
v := Vertex{1, 2}

// Field Access with .
v.X = 4

// Struct Pointer
p := &v
p.X = 4  // No need to dereference (*p).X
```

## Arrays
- Cannot be resized. Size is part of its type.
```
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Slices
- Dynamically sized arrays.

- Type: `[]T`
- Python Slices - `a[low:high]`
- Slices does not store data. It is a reference to an underlying array. Other slices that share the same underlying array will see those changes.
- Length `len(s)` - number of elements in the slice
- Capacity `cap(s)` - number of elements in the underlying array
- Zero Value of slice is `nil`
- Create a slice with `make([]int, lenValue, capValue)`
- 2D slice (see [dynamically allocated 2D slice](./exercise-slices.go))
```
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
}
```
- Append to a slice by using append `slice = append(slice, values...)`. A larger array will be reallocated if needed.

## Maps
- Zero value is `nil`, keys cannot be added
```
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// Add elem
m[key] = elem

// Get elem
elem = m[key]

// Delete elem
delete(m, key)

// Test if key is present
// ok = true if key is present, else false
elem, ok = m[key]
```
- If top-level type is specified, can be omitted when specifiying key/values

- See [this example](./exercise-maps.go).

## Function Values
- Functions can be arguments and return values.

- See [this example](./function-values.go).
- Closures is a function that references values outside of its body.
- Functions can be closures. [Example](./function-closures.go) [Example2](./exercise-fibonacci-closure.go)

