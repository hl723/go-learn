# A Tour of Go - Methods and Interfaces

## Methods
- Just a function with a receiver argument (the object)

- Can only define methods for types in the *same* package. Could do something like `type MyFloat float64` for primitive types.

- Pointer receivers are more common than value receivers.

- Pointer receiver type cannot already be a pointer.

- Value Receiver - pass by value (a copy)

- Pointer Receiver - pass by reference

- When invoking methods, pointer dereferencing is automatically done. 

- All methods of a type should either be value or pointer receivers (in general, due to interfaces)

## Interfaces
- Interface - set of method signatures.

```
type Abser interface {
	Abs() float64
}
```
- A value of interface type can hold any value that implement the methods.

- Interfaces are implemented implicitly. No "implements" keyword. Decouples interface with implementation. Can appear in any package.

- Interfaces are just a tuple under the hood. `(value, underlying type)`

- If the value is nil, this will not trigger a null pointer exception, but passed into the method to handle.

- If the interface value (the tuple) is itself nil, then this is a runtime error.

- Empty Interface `interface{}` is a interface with zero *specified* methods. These are used for code handling variables of unknown types.

- Type Assertions - Test for type of unknown variable
```
var i interface{} = "hello"

s := i.(string) // s ("hello")
fmt.Println(s)

s, ok := i.(string) // s ("hello"), ok (true)
fmt.Println(s, ok)

f, ok := i.(float64) // f (0), ok (false) error is captured
fmt.Println(f, ok)

f = i.(float64) // panic, as we don't catch the error
fmt.Println(f)
```
- Type Checking in Series with Switch Statement
```
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

## Stringers
- Stringer - similar to Java's toString() method, this is the method interface when fmt looks for to print values. [Example](./exercise-stringer.go) 
```
type Stringer interface {
    String() string
}
```

## Error Type
- The `error` type is a built-in interface type. Looks for this interface when printing error messages. [Example](./excercise-errors.go)
```
type error interface {
    Error() string
}
```

## Readers
- The `io.Reader` interface is as follows.
```
func (T) Read(b []byte) (n int, err error)
```
- It populates the byte slice with data. Returning n (num bytes read) and an error value.
- The Reader will return `io.EOF` error when the stream ends. 
```
// Example 1
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// A reader that emits an infinite stream of ASCII character 'A's
func (r MyReader) Read(buffer []byte) (n int, err error) {
	n = 0
	err = nil // never return io.EOF since it's an infinite stream
	for i, _ := range buffer {
		buffer[i] = byte('A')
		n += 1	
	}
	return
}

// Example 2 Rot13Reader Wrapper
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	if n > 0 {
		for i := range n {
			if b[i] >= byte('A') && b[i] <= byte('Z') {
				b[i] = ((b[i] - byte('A') + 13) % 26) + byte('A')
			} else if b[i] >= byte('a') && b[i] <= byte('z') {
				b[i] = ((b[i] - byte('a') + 13) % 26) + byte('a')
			}
		}
	}
	return n, err
}
```

## Images
```
// Image Interface
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
- [Example](./exercise-images.go)
