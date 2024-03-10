# A Tour of Go - Generics

- Generic types are defined in square brackets.
```
func Index[T comparable](s []T, x T) int
```

- `comparable` is a constraint for types with `==` and `!=` operators on the type.

- See [this example](../generics/go.mod).
