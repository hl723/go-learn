# Format Specifiers in Go

This is a cheatsheet of all the format specifiers for fmt.Printf and fmt.Sprintf.

## Default Formats and Type
- %v - default format
- %#v - Go-syntax format
- %T - type of variable

## Integers
- %d - base 10 integer
- %+d - show sign (+/-)
- %4d - pad with spaces, width 4, right-justified
- %-4d - same but left-justified
- %04d - pad with zeros
- %b - binary
- %o - octal
- %x - hex lowercase
- %X - hex uppercase
- %#x - hex with leading 0x

## Floats
- %e - scientific notation
- %f - floating point
- %.2f - default width, 2 precision
- %8.2f - 8 width, 2 precision
- %g - Exponent as needed

## Character
- %c - char
- %q - quoted char
- %U - unicode value of char
- %#U - unicode value AND char in quotes

## Strings or Byte Slice
- %s - string
- %q - quoted string
- %x - hex dump of bytes (no spaces)
- % x - hex dump of bytes (with spaces)

## Boolean
- %t - prints "true"/"false"

## Pointer
- %p - pointer in hex with 0x
